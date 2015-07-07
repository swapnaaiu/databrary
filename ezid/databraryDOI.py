import ezid_api
import logging 
import psycopg2
from lxml import etree as e
from config import conn as c #TODO: THIS WILL HAVE TO CHANGE
import datacite_validator as dv
import datetime
import sys, os

#check if test or not - pass any argument for it to be a test run
if len(sys.argv)>1:    
    istest = sys.argv[1]
else:
    istest = None
#set globals
LOG_PATH = './logs/'
LOG_FILE = 'ezidlog.log'
LOG_DEST = LOG_PATH + LOG_FILE
MAX_LOG_SIZE = 10000000
DOI_1 = 'doi:10.17910/B7159Q'
TEST = False
if istest is not None:
    TEST = True

def __setup_log():
    '''check if the log file exists or not, if not, create if, if so 
       check if it's at 10mb or less. If more, we want to create a new one
       and move the log file to a stored version'''
    if not os.path.exists(LOG_PATH):
        os.mkdir(LOG_PATH)
    if not os.path.isfile(LOG_DEST):
        logfile = open(LOG_DEST, 'w+') 
    else: 
        stats = os.stat(LOG_DEST)
        size = stats.st_size
        if size > MAX_LOG_SIZE:
            existing_logs = os.listdir(LOG_PATH)
            if len(existing_logs) > 1:
                copies = [int(i.split('_')[1].split('.')[0]) for i in existing_logs if '_' in i]
                increment = max(copies) + 1
                newdest = LOG_PATH + 'ezidlog_' + str(increment) + '.log' 
                os.rename(LOG_DEST, newdest)
                logfile = open(LOG_DEST, 'w+')
            else:    
                newdest = LOG_PATH + 'ezidlog_0.log'
                os.rename(LOG_DEST, newdest)
                logfile = open(LOG_DEST, 'a')
        else:
            logfile = open(LOG_DEST, 'a') 
    return logfile
__setup_log()

#Initiate and configure the logger
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)
handler = logging.FileHandler(LOG_DEST)
handler.setLevel(logging.INFO)
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)

target_path = "http://databrary.org"

sql = { 'QueryAll' : ("SELECT v.id as target, volume_creation(v.id), volume_access_check(v.id, -1) > 'NONE', v.name as title, "
            "COALESCE(sortname || ', ' || prename, sortname, '') as creator, p.id as party_id, volume_doi.doi, volume_citation.*, v.body "
            "FROM volume v "
            "LEFT JOIN volume_access va1 ON v.id = va1.volume AND va1.individual = 'ADMIN' "
            "JOIN party p ON p.id = va1.party "
            "LEFT JOIN volume_doi ON v.id = volume_doi.volume "
            "LEFT JOIN volume_citation ON v.id = volume_citation.volume "
            "ORDER BY target;"
            ), 
       'GetFunders' : "SELECT vf.volume, vf.awards, f.name, f.fundref_id FROM volume_funding vf LEFT JOIN funder f ON vf.funder = f.fundref_id WHERE volume IN %s;", 
       'AddDOI' : "INSERT into volume_doi VALUES (%s, %s)"}

class dbDB(object):
    _conn = None
    _cur = None

    def __init__(self):
        try:
            self._conn = psycopg2.connect(dbname=c._CREDENTIALS['db'],
                                          user=c._CREDENTIALS['u'],
                                          host=c._CREDENTIALS['host'],
                                          password=c._CREDENTIALS['p'])
        except Exception as e:
            logger.error("Unable to connect to database. Exception: %s. Script coming to a screeching halt" % str(e))
        self._cur = self._conn.cursor()

    def __del__(self):
        return self._conn.close()

    def query(self, query, params=None):
        return self._cur.execute(query, params)

    def insert(self, insert, params=None):
        return self._cur.execute(insert, params)
       

def queryAll(db):
    try:
        db.query(sql['QueryAll'])
        return db._cur.fetchall()
    except Exception as e:
        logger.error("Unable to query database. Exception: %s. Script coming to a screeching halt" % str(e))

def _getFunders(db, vs): #vs is a tuple of volume ids -> dict
    try:
        db.query(sql['GetFunders'], (vs,))
        funders = db._cur.fetchall()
        funder_data = {f[0]:[] for f in funders}
        for f in funders:
            funder_data[f[0]].append({"award_no":f[1], "funder":f[2], "fundref_id":f[3]})
        return funder_data
    except Exception as e:
        logger.error("Unable to query database. Exception: %s. Script coming to a screeching halt" % str(e))

def _addDOIS(db, new_dois):
    '''takes a list of dicts with dois and the volumes they belong to and updates the database with those values'''

    try:
        for i in new_dois:
            params = (i['vol'], i['doi'])
            db.insert(sql['AddDOI'], params)
        db._conn.commit()
    except Exception as e:
            logger.error("Unable to insert to database. Exception: %s. Script coming to a screeching halt" % str(e))

def _createXMLDoc(row, volume, creators, funders, doi=None): #tuple, str, dict, dict, dict, str, -> xml str
    '''taking in a row returned from the database, convert it to datacite xml
        according to http://ezid.cdlib.org/doc/apidoc.html#metadata-profiles this can
        be then sent along in the ANVL'''
    vol_date = row[1]
    vol_title = row[3].decode('utf-8') if row[3] is not None else "(:unav)"
    cite_url = row[9]
    vol_body = row[11].decode('utf-8') if row[11] is not None else "(:unav)"
    xmlns="http://datacite.org/schema/kernel-3"
    xsi="http://www.w3.org/2001/XMLSchema-instance"
    schemaLocation="http://datacite.org/schema/kernel-3 http://schema.datacite.org/meta/kernel-3/metadata.xsd"
    fundrefURI = "http://data.fundref.org/fundref/funder/"
    NSMAP = {None:xmlns, "xsi":xsi}
    xmldoc = e.Element("resource", attrib={"{"+xsi+"}schemaLocation":schemaLocation},  nsmap=NSMAP)
    ielem = e.SubElement(xmldoc, "identifier", identifierType="DOI")
    ielem.text = "(:tba)" if doi is None else doi
    pubelem = e.SubElement(xmldoc, "publisher")
    pubelem.text = "Databrary"
    pubyrelem = e.SubElement(xmldoc, "publicationYear")
    pubyrelem.text = str(vol_date.year) if vol_date is not None else "(:unav)"
    telem = e.SubElement(xmldoc, "titles")
    title = e.SubElement(telem, "title")
    title.text = vol_title
    reselem = e.SubElement(xmldoc, "resourceType", resourceTypeGeneral="Dataset")
    reselem.text = "Dataset"
    descelem = e.SubElement(xmldoc, "descriptions")
    description = e.SubElement(descelem, "description", descriptionType="Abstract")
    description.text = vol_body if vol_body is not None else "(:unav)"
    crelem = e.SubElement(xmldoc, "creators")
    felem = e.SubElement(xmldoc, "contributors")
    for c in creators[volume]:
        cr = e.SubElement(crelem, "creator")
        crname = e.SubElement(cr, "creatorName")
        crname.text = c.decode('utf-8')
    if volume in funders.keys():
        for f in funders[volume]:   
            ftype = e.SubElement(felem, "contributor", contributorType="Funder")
            fname = e.SubElement(ftype, "contributorName")
            fname.text = f['funder'].decode('utf-8')
            fid = e.SubElement(ftype, "nameIdentifier", schemeURI=fundrefURI, nameIdentifierScheme="FundRef")
            fid.text = fundrefURI + str(f['fundref_id'])
    if cite_url is not None:
        if cite_url.startswith('doi'):
            cite_url = "http://dx.doi.org/" + cite_url.split(':')[1]
        relelem = e.SubElement(xmldoc, "relatedIdentifiers")
        relid = e.SubElement(relelem, "relatedIdentifier", relatedIdentifierType="URL", relationType="IsSupplementTo")
        relid.text = cite_url
    xmloutput = e.tostring(xmldoc)
    #dv._validateXML(xmloutput) # this is not ideal because we have to send a (:tba) to ezid, even though that doesn't validate
    return xmloutput

def _getCreators(rs): #rs is a list of rows
    '''compile all admin for a volume into a list of creators per volume'''
    creators = {r[0]:[] for r in rs}
    for r in rs:
        creators[r[0]].append(r[4])
    return creators

def _generateRecord(status, xml, vol):
    target_base = target_path + "/volume/"
    record = {"_target": target_base + str(vol), 
                             "_profile": "datacite", 
                             "_status": status, 
                             "datacite": xml
                            }
    return record

def _payloadDedupe(records, record_kind):
    '''dedupe records (since there's one row for every owner on the volume)'''
    set_list = []
    for m in records[record_kind]:
        if m not in set_list:
            set_list.append(m)
    return set_list

def makeMetadata(db, rs): #rs is a list -> list of metadata dict
    allToUpload = {"mint":[], "modify":[]}
    volumes = tuple(set([r[0] for r in rs]))
    funders = _getFunders(db, volumes)
    creators = _getCreators(rs)
    for r in rs:
        vol = r[0]
        shared = r[2]
        vol_doi = r[6]
        if shared is True and vol_doi is None:
            status = "public"
            xml = _createXMLDoc(r, vol, creators, funders)
            allToUpload['mint'].append({"volume":vol, "record":_generateRecord(status, xml, vol)})
        elif shared is not True and vol_doi is not None:
            status = "unavailable"
            xml = _createXMLDoc(r, vol, creators, funders, vol_doi)
            allToUpload['modify'].append({'_id':"doi:"+vol_doi, 'record':_generateRecord(status, xml, vol)})
        elif shared is True and vol_doi is not None:
            status = "public"
            xml = _createXMLDoc(r, vol, creators, funders, vol_doi)
            allToUpload['modify'].append({'_id':"doi:"+vol_doi, 'record':_generateRecord(status, xml, vol)})

    mdPayload = {'mint':_payloadDedupe(allToUpload, 'mint'), "modify":_payloadDedupe(allToUpload, 'modify')}
    return mdPayload

def postData(db, payload):
    new_dois = []
    ezid_doi_session = ezid_api.ApiSession(username=c._CREDENTIALS['ezid_u'], password=c._CREDENTIALS['ezid_p'], scheme='doi')
    #check if the server is up, if not, bail
    server_response = ezid_doi_session.checkserver()
    if server_response == True:
    	logger.info('ezid server is up, starting operation')
    else:
    	logger.info('ezid server seems to be down, exiting since will not be able to do anything until it comes back up')
    	sys.exit()
    logger.info('minting %s DOIs and modifiying %s existing records' % (str(len(payload['mint'])), str(len(payload['modify']))))
    #start by minting any new shares
    for p in payload['mint']:
        volume = p['volume']
        record = p['record']
        if TEST == True:
            mint_res = "Your DOI for %s will not be minted because this is test mode" % volume
            print mint_res
        else:
            mint_res = ezid_doi_session.mint(record)
        if mint_res.startswith('doi'):
            curr_doi = mint_res.split('|')[0].strip().split(':')[1]
            new_dois.append({'vol':volume, 'doi':curr_doi})
            logger.info('minted doi: %s for volume %s' % (curr_doi, volume))
        else:
            logger.error('something went wrong minting a DOI for volume %s, error returned: %s' % (volume, mint_res))
    _addDOIS(db, new_dois)
    #next update existing records with dois
    for q in payload['modify']:
        identifier = q['_id']
        record = q['record']
        new_status = record['_status']
        mod_res = ezid_doi_session.recordModify(identifier, record)
        if type(mod_res) == dict:
            logger.info('%s successfully modified' % identifier)
        elif mod_res.startswith('error'):
            logger.error('something went wrong modifying a record for DOI: %s, error returned: %s' % (identifier, mod_res))

if __name__ == "__main__":
    db = dbDB()
    rows = queryAll(db)
    tosend = makeMetadata(db, rows)
    postData(db, tosend)
    del db