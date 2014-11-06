'use strict'

app.directive 'volumeEditReferencesForm', [
  'pageService',
  (page) ->
    restrict: 'E',
    templateUrl: 'volume/editReferences.html',
    link: ($scope) ->
      volume = $scope.volume
      form = $scope.volumeEditReferencesForm

      form.data = volume.references.map (ref) ->
        head: ref.head
        url: ref.url
      form.data.push
        head: ''
        url: ''

      form.change = () ->
        if form.data[form.data.length-1].url != ''
          form.data.push
            head: ''
            url: ''
        else if form.data.length == 1
          # hack!
          page.$timeout () ->
            form.$setDirty()

      form.save = () ->
        volume.save({references: form.data}).then(() ->
            form.validator.server {}

            form.messages.add
              type: 'green'
              countdown: 3000
              body: page.constants.message('volume.edit.success')

            form.$setPristine()
          , (res) ->
            form.validator.server res
          )

      form.scrollFn = page.display.makeFloatScrollFn($('.ver-float'), $('.ver-float-floater'), 24*2.5)
      page.$w.scroll(form.scrollFn)
]
