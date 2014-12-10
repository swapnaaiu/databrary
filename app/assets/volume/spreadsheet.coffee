'use strict'

app.directive 'spreadsheet', [
  'constantService', 'displayService', 'messageService', 'tooltipService', '$compile', '$templateCache', '$timeout',
  (constants, display, messages, tooltips, $compile, $templateCache, $timeout) ->
    maybeInt = (s) ->
      if isNaN(i = parseInt(s, 10)) then s else i
    byNumber = (a,b) -> a-b
    byId = (a,b) -> a.id-b.id
    byType = (a,b) ->
      ta = typeof a
      tb = typeof b
      if ta != tb
        a = ta
        b = tb
      if a>b then 1 else if a<b then -1 else 0
    byMagic = (a,b) ->
      na = parseFloat(a)
      nb = parseFloat(b)
      return 1 if na > nb
      return -1 if na < nb
      byType(a,b)

    stripPrefix = (s, prefix) ->
      if s.startsWith(prefix) then s.substr(prefix.length)

    # autovivification
    arr = (a, f) ->
      if f of a then a[f] else a[f] = []
    obj = (a, f) ->
      if f of a then a[f] else a[f] = {}
    inc = (a, f) ->
      if f of a then a[f]++ else
        a[f] = 1
        0

    parseInfo = (id) ->
      return if id == undefined
      s = id.split '_'
      info = {t: s[0]}
      return if s.length > 1 && isNaN(info.i = parseInt(s[1], 10))
      switch info.t
        when 'rec'
          if 3 of s
            info.n = parseInt(s[2], 10)
            info.m = parseInt(s[3], 10)
          else
            info.n = 0
            info.m = parseInt(s[2], 10)
        when 'add', 'more'
          info.c = parseInt(s[2], 10)
        when 'metric'
          info.m = info.i
          delete info.i
        when 'category'
          info.c = info.i
          delete info.i
      info

    pseudoCategory =
      0:
        id: 0
        name: 'record'
        not: 'No record'
        template: [constants.metricName.ident.id]
      asset:
        id: 'asset'
        name: 'file'
        not: 'No file'
    constants.deepFreeze(pseudoCategory)
    getCategory = (c) ->
      pseudoCategory[c || 0] || constants.category[c]

    pseudoMetric =
      id:
        id: 'id'
        name: 'id'
        display: ' '
        type: 'number'
        classification: constants.classification.PUBLIC
      age:
        id: 'age'
        name: 'age'
        type: 'number'
        classification: constants.classification.SHARED
    constants.deepFreeze(pseudoMetric)
    getMetric = (m) ->
      pseudoMetric[m] || constants.metric[m]

    selectStyles = document.head.appendChild(document.createElement('style')).sheet

    {
    restrict: 'E'
    scope: true
    templateUrl: 'volume/spreadsheet.html'
    controller: [
      '$scope', '$element', '$attrs',
      ($scope, $element, $attrs) ->
        volume = $scope.volume

        editing = $scope.editing = $attrs.edit != undefined
        top = $scope.top = 'top' of $attrs
        assets = $scope.assets = 'assets' of $attrs
        id = $scope.id = $attrs.id ? if top then 'sst' else 'ss'

        ###
        # We use the following types of data structures:
        #   Row = index of slot in slots and rows (i)
        #   Data[Row] = scalar value (array over Row)
        #   Slot_id = Database id of container
        #   Segment = standard time range (see type service)
        #   Record_id = Database id of record
        #   Category_id = Database id of record category (c)
        #   Count = index of record within category for slot (n)
        #   Metric_id = Database id of metric, or "id" for Record_id, or "age" (m)
        ###

        ### jshint ignore:start #### fixed in jshint 2.5.7
        slots = (container for containerId, container of volume.containers when top != !container.top) # [Row] = Slot
        ### jshint ignore:end ###

        order = Object.keys(slots)  # Permutation Array of Row in display order

        records = {}                # [Category_id][Metric_id][Count] :: Data
        counts = new Array(slots.length) # [Row][Category_id] :: Count
        recordCols = []             # [] Array over records :: {category: Category_id, metrics[]: Array of Metric_id}
        metricCols = []             # [] Array over metrics :: {category: Category, metric: Metric} (flattened version of recordCols)
        depends = {}                # [Record_id][Row] :: Count

        rows = new Array(slots.length) # [Row] :: DOM Element tr

        tbody = $element[0].getElementsByTagName("tbody")[0]

        parseId = (el) ->
          info = parseInfo(stripPrefix(el.id, id+'-'))
          return info unless info
          if 'm' of info
            info.c = (info.category = (info.col = metricCols[info.m]).category).id
            info.metric = info.col.metric
          else if 'c' of info
            info.category = getCategory(info.c)
          if 'i' of info
            info.slot = slots[info.i]
            if 'n' of info
              info.record = volume.records[info.r = records[info.c].id[info.n][info.i]]
          info

        ################################# Populate data structures 

        # Fill all Data values for Row i
        populateSlot = (i) ->
          slot = slots[i]

          # r, n
          populateMeasure = (m, v) ->
            arr(arr(r, m), n)[i] = v
            return
          count = counts[i] = {}

          for rr in slot.records
            record = rr.record
            # temporary workaround for half-built volume inclusions:
            continue unless record
            c = record.category || 0

            # populate depends:
            if record.id of depends
              # skip duplicates:
              continue if i of depends[record.id]
            else
              depends[record.id] = {}

            # populate records:
            r = if c of records then records[c] else records[c] = {id: []}

            # determine Count:
            n = inc(count, c)

            # populate measures:
            populateMeasure('id', record.id)
            if !editing && 'age' of rr
              populateMeasure('age', rr.age)
            for m, v of record.measures
              populateMeasure(m, v)

            depends[record.id][i] = n

          count.asset = slot.assets.length if assets

          return

        # Fill metricCols and recordCols from records
        populateCols = ->
          metricCols = []
          $scope.recordCols = recordCols = Object.keys(records).sort(byNumber).map (c) ->
            category = getCategory(c)
            if editing
              for m in category.template
                arr(records[c], m)
            metrics = Object.keys(records[c]).map(maybeInt).sort(byType)
            metrics.pop() # remove 'id' (necessarily last)
            metrics = metrics.map(getMetric)
            # add back the 'id' column first if needed
            if !metrics.length || editing && !(metrics.length == 1 && metrics[0].options)
              metrics.unshift(pseudoMetric.id)
            si = metricCols.length
            metricCols.push.apply metricCols, metrics.map (m) ->
              category: category
              metric: m
              sortable: m != pseudoMetric.id || metrics.length == 1
            l = metrics.length
            metricCols[si].first = metricCols[si+l-1].last = l
            {
              category: category
              metrics: metrics
              start: si
            }
          $scope.metricCols = metricCols
          return

        # Call all populate functions
        populate = ->
          records = {}
          depends = {}
          for s, i in slots
            populateSlot(i)
          populateCols()
          generate()
          return

        ################################# Generate HTML
        
        # Find the text content of cell c with element t
        setCell = (c, t) ->
          el = c.lastChild
          if el && el.nodeType == 3
            c.replaceChild(t, el)
          else
            c.appendChild(t)
          return

        # Add or replace the text contents of cell c for measure/type m with value v
        generateText = (c, m, v, assumed) ->
          if m == 'consent'
            cn = constants.consent[v || 0]
            c.className = cn + ' consent icon hint-consent-' + cn
            v = ''
          else if v == undefined
            c.classList.add('blank')
            v = assumed || ''
          else if m == 'classification' || m == 'excerpt'
            cn = constants.classification[v]
            c.className = cn + ' classification icon hint-classification-' + cn
            v = ''
          else
            c.classList.remove('blank')
            if m == 'id'
              c.className = 'icon ' + if editing then 'trash' else 'bullet'
              v = ''
            else if m == 'age'
              v = display.formatAge(v)
          setCell(c, document.createTextNode(v))
          return

        # Add a td element to tr r with value c and id i
        generateCell = (r, m, v, i, assumed) ->
          td = r.appendChild(document.createElement('td'))
          if v == null
            td.className = 'null'
          else
            generateText(td, m, v, assumed)
            td.id = i
          td

        generateMultiple = (cat, cols, row, i, n, t) ->
          if n == undefined
            return if t == 1
          else
            return if n < t
          td = row.appendChild(document.createElement('td'))
          td.setAttribute("colspan", cols)
          if n == undefined && t > 1
            td.appendChild(document.createTextNode(t + " " + cat.name + "s"))
            td.className = 'more'
            td.id = id + '-more_' + i + '_' + cat.id
          else
            td.className = 'null'
            if !n || n == t
              if !n
                td.appendChild(document.createTextNode(cat.not))
              else if editing
                td.appendChild(document.createTextNode("add " + cat.name))
              if edit
                td.className = 'null add'
                td.id = id + '-add_' + i + '_' + cat.id
          td

        # Add all the measure tds to row i for count n, record r
        generateRecord = (row, i, col, n) ->
          return unless l = col.metrics.length
          c = col.category.id
          t = counts[i][c] || 0
          r = records[c]
          if td = generateMultiple(col.category, l, row, i, n, t)
            if n == undefined
              for n in [0..t-1] by 1
                td.classList.add('ss-rec_' + r.id[n][i])
            return
          ms = col.metrics
          b = id + '-rec_' + i + '_'
          if n == undefined
            n = 0
          else
            b += n + '_'
          for mi in [0..l-1] by 1
            m = ms[mi].id
            v = r[m][n] && r[m][n][i]
            cell = generateCell(row, m, v, b + (col.start+mi), ms[mi].assumed)
            if v != null
              ri = 'ss-rec_' + r.id[n][i]
              cell.classList.add(ri)
              cell.classList.add(ri + '_' + m)
          return

        generateAsset = (row, i, n) ->
          a = slots[i].assets
          return if generateMultiple(pseudoCategory.asset, 3, row, i, n, a.length)
          a = a[n || 0]
          generateCell(row, 'asset', a.name, id + '-asset_' + a.id)
          generateCell(row, 'classification', a.classification, id + '-class_' + a.id)
          generateCell(row, 'excerpt', a.excerpt, id + '-excerpt_' + a.id)
          return

        # Fill out rows[i].
        generateRow = (i) ->
          slot = slots[i]
          stop = slot.id == volume.top.id
          row = document.createElement('tr')
          rows[i]?.parentNode?.replaceChild(row, rows[i])
          rows[i] = row
          row.id = id + '_' + i
          row.data = i
          if editing && stop
            row.className = 'top'

          name = slot.name
          if stop
            name ?= constants.message('materials.top')
          cell = generateCell(row, 'name', name, id + '-name_' + i)
          if editing && !stop
            a = cell.insertBefore(document.createElement('a'), cell.firstChild)
            a.className = 'trash icon'
            $(a).on 'click', (event) ->
              $scope.$apply () ->
                removeSlot(cell, i, slot)
                return
              event.stopPropagation()
              return
          a = cell.insertBefore(document.createElement('a'), cell.firstChild)
          a.setAttribute('href', if editing then slot.editRoute() else slot.route())
          a.className = "session icon hint-object-play"

          generateCell(row, 'date', slot.date, id + '-date_' + i) unless slot.top
          generateCell(row, 'consent', slot.consent, id + '-consent_' + i)
          for c in recordCols
            generateRecord(row, i, c)
          if assets
            generateAsset(row, i)
          return

        # Update all age displays.
        $scope.$on 'displayService-toggleAge', ->
          for m, mi in metricCols
            continue unless m.metric.id == 'age'
            c = m.category.id
            r = records[c][m.metric.id]
            if expandedCat == c && counts[expanded][c] > 1
              for n in [0..counts[expanded][c]-1] by 1 when n of r
                generateText(
                  document.getElementById(id + '-rec_' + expanded + '_' + n + '_' + mi),
                  'age', r[n][expanded])
            return unless 0 of r
            r = r[0]
            post = '_' + mi
            for d, i in r when counts[i][c] == 1
              generateText(
                document.getElementById(id + '-rec_' + i + post),
                'age', d)

        # Generate all rows.
        generate = ->
          for s, i in slots
            generateRow(i)
          fill()
          return

        ################################# Place DOM elements
        
        # Place all rows into spreadsheet.
        fill = ->
          collapse()
          for i in order
            tbody.appendChild(rows[i])
          return

        # Populate order based on compare function applied to values.
        sort = (values, compare) ->
          return unless values
          compare ?= byMagic
          order.sort (i, j) ->
            compare(values[i], values[j])
          return

        currentSort = undefined
        currentSortDirection = false
  
        # Sort by values, called name.
        sortBy = (key, values) ->
          if currentSort == key
            currentSortDirection = !currentSortDirection
            order.reverse()
          else
            sort(values)
            currentSort = key
            currentSortDirection = false
          fill()
          return

        # Sort by one of the container columns.
        sortBySlot = (f) ->
          sortBy(f, slots.map((s) -> s[f]))

        # Sort by Category_id c's Metric_id m
        sortByMetric = (col) ->
          sortBy(col, records[col.category.id][col.metric.id][0])

        $scope.colClasses = (col) ->
          cls = []
          if typeof col == 'object'
            cls.push 'first' if col.first
            cls.push 'last' if col.last
            cls.push 'sort' if col.sortable
          else
            cls.push 'sort'
          if currentSort == col
            cls.push 'sort-'+(if currentSortDirection then 'desc' else 'asc')
          else
            cls.push 'sortable'
          cls

        ################################# Backend saving

        saveRun = (cell, run) ->
          cell.classList.remove('error')
          cell.classList.add('saving')
          run.then () ->
              cell.classList.remove('saving')
              return
            , (res) ->
              cell.classList.remove('saving')
              cell.classList.add('error')
              messages.addError
                body: 'error' # FIXME
                report: res
              return

        createSlot = (cell) ->
          saveRun cell, volume.createContainer({top:top}).then (slot) ->
            arr(slot, 'records')
            i = slots.push(slot)-1
            order.push(i)
            populateSlot(i)
            generateRow(i)
            tbody.appendChild(rows[i])
            return

        saveSlot = (cell, info, v) ->
          data = {}
          data[info.t] = v ? ''
          return if info.slot[info.t] == data[info.t]
          saveRun cell, info.slot.save(data).then () ->
            generateText(cell, info.t, info.slot[info.t])
            return

        removeSlot = (cell, i, slot) ->
          # assuming we have a container
          saveRun cell, slot.remove().then (done) ->
            unless done
              messages.add
                body: constants.message('slot.remove.notempty')
                type: 'red'
                countdown: 5000
              return
            unedit()
            collapse()
            rows[i].parentNode.removeChild(rows[i])
            slots.splice(i, 1)
            counts.splice(i, 1)
            rows.splice(i, 1)
            order.remove(i)
            order = order.map (j) -> j - (j > i)
            populate()
            return

        saveMeasure = (cell, record, metric, v) ->
          return if record.measures[metric.id] == v
          saveRun cell, record.measureSet(metric.id, v).then (rec) ->
            rcm = records[rec.category || 0][metric.id]
            for i, n of depends[record.id]
              arr(rcm, n)[i] = v
              # TODO age may have changed... not clear how to update.
            l = tbody.getElementsByClassName('ss-rec_' + record.id + '_' + metric.id)
            for li in l
              generateText(li, metric.id, v, metric.assumed)
            return

        setRecord = (cell, info, record) ->
          add = ->
            if record
              info.slot.addRecord(record)
            else if record != null
              info.slot.newRecord(info.c || '')
          act =
            if info.record
              info.slot.removeRecord(info.record).then(add)
            else
              add()

          saveRun cell, act.then (record) ->
            if record
              r = record.id
              info.n = inc(counts[info.i], info.c) unless 'n' of info

              for m, rcm of records[info.c]
                v = if m of record then record[m] else record.measures[m]
                if v == undefined
                  delete rcm[info.n][info.i] if info.n of rcm
                else
                  arr(rcm, info.n)[info.i] = v
              # TODO this may necessitate regenerating column headers
            else
              t = --counts[info.i][info.c]
              for m, rcm of records[info.c]
                for n in [info.n+1..rcm.length-1] by 1
                  arr(rcm, n-1)[info.i] = arr(rcm, n)[info.i]
                delete rcm[t][info.i] if t of rcm

            delete depends[info.r][info.i] if info.record
            obj(depends, r)[info.i] = info.n if record

            collapse()
            generateRow(info.i)
            expand(info) if info.n
            record

        ################################# Interaction

        expandedCat = undefined
        expanded = undefined

        # Collapse any expanded row.
        collapse = ->
          return if expanded == undefined
          i = expanded
          expanded = expandedCat = undefined
          row = rows[i]
          row.classList.remove('expand')
          t = 0
          while (el = row.nextSibling) && el.data == i
            t++
            tbody.removeChild(el)

          el = row.firstChild
          while el
            el.removeAttribute("rowspan")
            el = el.nextSibling

          t

        # Expand (or collapse) a row
        expand = (info) ->
          if expanded == info.i && expandedCat == info.c
            if info.t == 'more'
              collapse()
            return
          collapse()

          expanded = info.i
          expandedCat = info.c
          row = rows[expanded]
          row.classList.add('expand')

          max = counts[expanded][expandedCat]
          max++ if editing
          return if max <= 1
          next = row.nextSibling
          start = counts[expanded][expandedCat] == 1
          col = expandedCat != 'asset' &&
            recordCols.find (col) -> col.category.id == expandedCat
          for n in [+start..max-1] by 1
            el = tbody.insertBefore(document.createElement('tr'), next)
            el.data = expanded
            el.className = 'expand'
            if col
              generateRecord(el, expanded, col, n)
            else
              generateAsset(el, expanded, n)

          max++ unless start
          el = row.firstChild
          while el
            info = parseId(el)
            if !info || info.c != expandedCat
              el.setAttribute("rowspan", max)
            el = el.nextSibling
          return

        save = (cell, type, value) ->
          info = parseId(cell)
          if value == ''
            value = undefined
          else switch type
            when 'consent'
              value = parseInt(value, 10)
            when 'record'
              if value == 'new'
                setRecord(cell, info)
              else if value == 'remove'
                setRecord(cell, info, null)
              else if v = stripPrefix(value, 'add_')
                u = v.indexOf('_')
                m = constants.metric[v.slice(0,u)]
                v = v.slice(u+1)
                setRecord(cell, info).then (r) ->
                  saveMeasure(cell, r, m, v) if r
                  return
              else if !isNaN(v = parseInt(value, 10))
                if v != info.r
                  setRecord(cell, info, volume.records[v])
                else
                  edit(cell, info, true)
              return
            when 'metric'
              if value != undefined
                arr(records[info.c], value)
                populateCols()
                generate()
              return
            when 'category'
              if value != undefined
                arr(obj(records, value), 'id')
                populateCols()
                generate()
              return

          switch info.t
            when 'name', 'date', 'consent'
              saveSlot(cell, info, value)
            when 'rec'
              saveMeasure(cell, info.record, info.metric, value)

        editScope = $scope.$new(true)
        editScope.constants = constants
        editInput = editScope.input = {}
        editCellTemplate = $compile($templateCache.get('volume/spreadsheetEditCell.html'))
        editCell = undefined

        unedit = (event) ->
          return unless edit = editCell
          editCell = undefined
          $(edit).children('[name=edit]').off()
          return unless cell = edit.parentNode
          cell.removeChild(edit)
          cell.classList.remove('editing')
          tooltips.clear()

          save(cell, editScope.type, editInput.value) if event
          return
        editScope.unedit = unedit

        edit = (cell, info, alt) ->
          return if info.slot?.id == volume.top.id
          switch info.t
            when 'name'
              editScope.type = 'text'
              editInput.value = info.slot.name
            when 'date'
              editScope.type = 'date'
              editInput.value = info.slot.date
            when 'consent'
              editScope.type = 'consent'
              editInput.value = (info.slot.consent || 0) + ''
            when 'rec', 'add'
              if info.t == 'rec' && info.metric.id == 'id'
                setRecord(cell, info, null)
                return
              if info.t == 'rec' && (!info.col.first || alt)
                m = info.metric.id
                # we need a real metric here:
                return unless typeof m == 'number'
                editInput.value = volume.records[info.r].measures[m] ? ''
                if info.metric.options
                  editScope.type = 'select'
                  editScope.options = [''].concat(info.metric.options)
                else if info.metric.long
                  editScope.type = 'long'
                else
                  editScope.type = info.metric.type
                break
            # when 'add', fall-through
              c = info.category
              if 'r' of info
                editInput.value = info.r + ''
              else
                editInput.value = 'remove'
              editScope.type = 'record'
              editScope.options =
                new: 'Create new ' + c.name
                remove: c.not
              for ri, r of volume.records
                if (r.category || 0) == c.id && (!(ri of depends && info.i of depends[ri]) || ri == editInput.value)
                  editScope.options[ri] = r.displayName
              # detect special cases: singleton or unitary records
              for mi of records[c.id]
                mm = constants.metric[mi]
                if !m
                  m = mm
                else if mm
                  m = null
                  break
              if m == undefined && Object.keys(editScope.options).length > 2
                # singleton: id only, existing record(s)
                delete editScope.options['new']
              else if m && m.options
                # unitary: single metric with options
                delete editScope.options['new']
                for o in m.options
                  for ri, r of volume.records
                    return if (r.category || 0) == c.id && r.measures[m.id] == o
                  editScope.options['add_'+m.id+'_'+o] = o
            when 'category'
              editScope.type = 'metric'
              editInput.value = undefined
              editScope.options = []
              for mi, m of constants.metric when !(mi of records[info.c])
                editScope.options.push(m)
              editScope.options.sort(byId)
            when 'head'
              editScope.type = 'category'
              editInput.value = undefined
              editScope.options = []
              for ci, c of constants.category when ci not of records
                editScope.options.push(c)
              editScope.options.sort(byId)
              editScope.options.push(pseudoCategory[0]) unless 0 of records
            else
              return

          e = editCellTemplate editScope, (e) ->
            cell.insertBefore(editCell = e[0], cell.firstChild)
            cell.classList.add('editing')
            return
          e.on 'click', ($event) ->
            # prevent other ng-click handlers from taking over
            $event.stopPropagation()
            return

          tooltips.clear()
          $timeout ->
            input = e.children('[name=edit]')
            input.focus()
            # chrome produces spurious change events on date fields, so we rely on key-enter instead.
            input.one('change', $scope.$lift(unedit)) unless editScope.type == 'date'
            return
          return

        unselect = ->
          while selectStyles.cssRules.length
            selectStyles.deleteRule(0)

          unedit()
          return

        $scope.$on '$destroy', unselect

        select = (cell, info) ->
          unselect()
          expand(info)
          if info.t == 'rec'
            for c, ci in cell.classList when c.startsWith('ss-rec_')
              selectStyles.insertRule('.' + c + '{background-color:' +
                (if c.contains('_', 7) then '#e8e47f' else 'rgba(242,238,100,0.4)') +
                ';\n text-}', selectStyles.cssRules.length)

          edit(cell, info) if editing
          return

        $scope.click = (event) ->
          el = event.target
          return unless el.tagName == 'TD'
          return unless info = parseId(el)

          select(el, info)
          if 'm' of info && metricCols[info.m].metric.id == 'age'
            display.toggleAge()
          return

        $scope.clickSlot = ($event, t) ->
          if t
            sortBySlot(t, $event)
          else
            unselect()
          return

        $scope.clickAdd = ($event) ->
          unselect()
          edit($event.target, {t:'head'})
          return

        $scope.clickCategoryAdd = ($event, col) ->
          unselect()
          edit($event.target.parentNode, {t:'category',c:col.category.id}) if editing
          return

        $scope.clickMetric = (col) ->
          sortByMetric(col) if col.sortable
          return

        $scope.clickNew = ($event) ->
          createSlot($event.target)
          return

        ################################# main

        $scope.refresh = ->
          unedit()
          collapse()
          populate()
          return

        populate()
        return
    ]
    }
]
