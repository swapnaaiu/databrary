package ingest

import scala.concurrent.Future
import macros._
import dbrary._
import models._

object Adolph extends Ingest {
  import Parse._

  private type ObjectParser[A] = Parser[A => A]
  private object ObjectParser {
    def apply[A](f : (A, String) => A) : ObjectParser[A] =
      Parser[A => A](s => f(_, s))
    def map[A,B](f : (A, B) => A, p : Parser[B]) : ObjectParser[A] =
      p.map(b => f(_, b))
    def option[A,B](f : (A, B) => A, p : Parser[B], nullif : String = "") : ObjectParser[A] =
      Parser[A => A] { s =>
	if (nullif != null && (s.isEmpty || s.equals(nullif)))
	  identity[A]
	else
	  f(_, p(s))
      }
  }

  private class CellParser[A](name : String, parse : ObjectParser[A])
    extends ColumnParser[A => A](name, parse)
  
  private def blankCellParser[A] : CellParser[A] =
    new CellParser[A]("<blank>", const[A => A](identity[A] _))

  private def parseCells[A](a : A, ps : Seq[CellParser[A]]) : ListParser[A] =
    ListParser[A] { l =>
      ps.foldLeft((a, l)) { (al, p) =>
	val (a, l) = al
	p.map(_(a))(l)
      }
    }

  private def measureParser[T](metric : Metric[T], p : Parser[T]) : Parser[MeasureV[T]] =
    Parser[MeasureV[T]](s => new MeasureV(metric, p(s)))

  private def dateMeasureParser(metric : Metric[Date]) =
    measureParser[Date](metric, date)

  private def measureMap(m : Measure[_]*) : Map[Int,Measure[_]] =
    Map(m.map(m => (m.metricId.unId, m)) : _*)

  abstract class Record(val category : RecordCategory) {
    protected val measures : Map[Int,Measure[_]]
    protected def idents : Iterable[Measure[_]]
    type Key
    def key : Key

    override def toString = category.name + ":" + idents.map(m => m.metric.name + ":" + m.datum).mkString(",")

    def getMeasure(m : Metric[_]) : Option[Measure[_]] =
      measures.get(m.id.unId)
    protected final def addMeasure(m : Measure[_]) : Map[Int,Measure[_]] = {
      val i = m.metricId.unId
      if (measures.contains(i))
	fail("duplicate measure for " + m.metric.name + " on " + this)
      measures.updated(i, m)
    }
    def withMeasure(m : Measure[_]) : Record

    def find(volume : Volume) : Future[Option[models.Record]] =
      for {
	l <- Record.findMeasures(volume, Some(category), idents.toSeq : _*)
	_ <- check(l.length <= 1,
	  PopulateException("multiple matching records for " + this, volume))
      }	yield (l.headOption)
    def populate(volume : Volume) : Future[models.Record] =
      for {
	ro <- find(volume)
	r <- Async.getOrElse(ro, Record.create(volume, Some(category)))
	_ <- Async.foreach[Measure[_], Unit](measures.values, { m =>
	  r.measures(m.metric).fold {
	    r.setMeasure(m).flatMap(check(_, 
	      PopulateException("failed to set measure for record " + this, r)))
	  } { c =>
	    check(c === m,
	      PopulateException("inconsistent mesaures for record " + this + ": " + m + " <> " + c, r))
	  }
	})
      } yield (r)
  }

  final case class IndicatorRecord(override val category : RecordCategory, measures : Map[Int,Measure[_]] = Map.empty) extends Record(category) {
    def idents = Nil
    type Key = Unit
    def key = ()
    def withMeasure(m : Measure[_]) = copy(measures = addMeasure(m))
  }

  final case class IdentRecord(override val category : RecordCategory, measures : Map[Int,Measure[_]] = Map.empty) extends Record(category) {
    def idents = getMeasure(Metric.Ident)
    type Key = String
    def key = getMeasure(Metric.Ident).fold(fail(category.name + " ident missing"))(_.datum)
    def withMeasure(m : Measure[_]) = copy(measures = addMeasure(m))
  }

  final case class Participant(measures : Map[Int,Measure[_]] = Map.empty)
    extends Record(RecordCategory.Participant) {
    def idents = getMeasure(Metric.Ident) ++ getMeasure(Metric.Info)
    type Key = (String, String)
    def key = (getMeasure(Metric.Ident).fold(fail("participant ID missing"))(_.datum), getMeasure(Metric.Info).fold("")(_.datum))
    def withMeasure(m : Measure[_]) = copy(measures = addMeasure(m))
  }

  private object Participants {
    private final val empty = Participant()
    private def measure[T](m : Parser[MeasureV[T]], nullif : String = "") : Parser[Participant => Participant] =
      ObjectParser.option[Participant,MeasureV[T]](_.withMeasure(_), m, nullif)
    val parseId = measure(measureParser(Metric.Ident, trimmed), null)
    val parseSet = measure(measureParser(Metric.Info, trimmed), null)
    private def parseHeader(name : String) : Parser[Participant => Participant] =
      name match {
	case "SUBJECT ID" => parseId
	case "DATASET" => parseSet
	case "BIRTH DATE" => measure(dateMeasureParser(Metric.Birthdate), null)
	case "GENDER" => measure(Gender.measureParse)
	case "RACE" => measure(Race.measureParse)
	case "ETHNICITY" => measure(Ethnicity.measureParse)
	case "TYPICAL DEVELOPMENT/DISABILITY" => measure(measureParser(Metric.Disability, trimmed), "typical")
	case _ => fail("unknown participant header: " + name)
      }
    private def header : Parser[CellParser[Participant]] =
      Parser[CellParser[Participant]] {
	case "" => blankCellParser
	case s => new CellParser(s, parseHeader(s))
      }
    private def parseData(l : List[List[String]]) : Map[Participant#Key,Participant] = l.zipWithIndex match {
      case h :: l =>
	val p = listAll(header).run(h)
	val line = parseCells(empty, p)
	l.foldLeft(Map.empty[Participant#Key,Participant]) { (m, l) =>
	  val p = line.run(l)
	  val i = p.key
	  if (m.contains(i))
	    throw ParseException("duplicate participant key: " + i, line = l._2)
	  m.updated(i, p)
	}
      case Nil => Map.empty[Participant#Key,Participant]
    }
    final def parseCSV(f : java.io.File) =
      parseData(CSV.parseFile(f))
  }

  final case class Exclusion(measures : Map[Int,Measure[_]] = Map.empty) extends Record(RecordCategory.Exclusion) {
    def idents = getMeasure(Metric.Reason)
    type Key = String
    def key = getMeasure(Metric.Reason).fold(fail("exclusion reason missing"))(_.datum)
    def withMeasure(m : Measure[_]) = copy(measures = addMeasure(m))
  }
  private object Exclusion {
    private object Reason extends MetricENUM(Metric.Reason) {
      val DID_NOT_MEET_INCLUSION_CRITERIA,
	PROCEDURAL_EXPERIMENTER_ERROR,
	FUSSY_TIRED_WITHDREW,
	OUTLIER = Value
    }

    def parse : Parser[Exclusion] =
      Reason.measureParse.map(m => Exclusion(measureMap(m)))
  }

  private def recordMap(r : Record*) : Map[Int,Record] =
    Map(r.map(r => (r.category.id.unId, r)) : _*)

  private final case class Session(name : String = "", date : Option[Date] = None, consent : Option[Consent.Value], records : Map[Int,Record] = recordMap(Participant())) {
    def withName(n : String) =
      if (name.nonEmpty && !name.equals(n))
	fail("duplicate session name: " + n)
      else copy(name = n)
    def withDate(d : Date) =
      if (!date.forall(_.equals(d)))
	fail("duplicate session date: " + d)
      else copy(date = Some(d))
    def withConsent(c : Consent.Value) =
      if (!date.forall(_.equals(c)))
	fail("duplicate session consent: " + c)
      else copy(consent = Some(c))
    def withParticipant(f : Participant => Participant) = {
      val i = RecordCategory.Participant.id.unId
      copy(records = records.updated(i, f(records(i).asInstanceOf[Participant])))
    }
    def withRecord(r : Record) = {
      val i = r.category.id.unId
      if (records.contains(i))
	fail("duplicate session record: " + r)
      copy(records = records.updated(i, r))
    }
  }

  private object Sessions {
    private def participant(p : Parser[Participant => Participant]) =
      ObjectParser.map[Session, Participant => Participant](_.withParticipant(_), p)
    private def parseHeader(name : String) : Parser[Session => Session] =
      name match {
	case "SUBJECT ID" => participant(Participants.parseId)
	case "DATASET" => participant(Participants.parseSet)
	case "TEST DATE" => ObjectParser.map[Session, Date](_.withDate(_), date)
	case "SESSION ID" => ObjectParser.option[Session, String](_.withName(_), trimmed)
	case "RELEASE LEVEL" => ObjectParser.option[Session, Consent.Value](_.withConsent(_), consent)
	case "PILOT" => ObjectParser.option[Session, Unit](
	  (s, _) => s.withRecord(IndicatorRecord(RecordCategory.Pilot)),
	  only("pilot"), "not pilot")
	case "EXCLUSION" => ObjectParser.option[Session, Exclusion](_.withRecord(_), Exclusion.parse, "not excluded")
	case _ => fail("unknown session header: " + name)
      }
    private def header : Parser[CellParser[Session]] =
      Parser[CellParser[Session]] {
	case "" => blankCellParser
	case s => new CellParser(s, parseHeader(s))
      }
  }

  def parseParticipants(f : java.io.File) : Iterable[Participant] =
    Participants.parseCSV(f).values
}
