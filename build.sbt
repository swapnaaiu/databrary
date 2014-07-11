name := "databrary"

scalaVersion in ThisBuild := "2.11.1"

scalacOptions in ThisBuild ++= Seq("-target:jvm-1.7","-optimise","-feature","-deprecation","-Xlint","-Yinline-warnings")

// scalacOptions += "-Ymacro-debug-lite"

resolvers in ThisBuild += Resolver.file("Local repo", file(Path.userHome.absolutePath+"/.ivy2/local"))(Resolver.ivyStylePatterns)

scalacOptions in (Compile, doc) <++= baseDirectory.map { bd => Seq(
  "-sourcepath", bd.getAbsolutePath,
  "-doc-source-url", "https://github.com/databrary/databrary/tree/master€{FILE_PATH}.scala"
) }

GitDescribe.gitDescribeOptions in ThisBuild := Seq("--long", "--dirty")

version in ThisBuild <<= GitDescribe.gitDescribe.apply(_.getOrElse("unknown"))

lazy val macros = project

lazy val dbrary = project
  .dependsOn(macros)

lazy val media = project
  .dependsOn(dbrary)

lazy val logbackAccess = project in file("logback-access")

lazy val databrary = (project in file("."))
  .enablePlugins(PlayScala, SbtWeb)
  .dependsOn(macros, dbrary, media, logbackAccess)

libraryDependencies ++= Seq(
  "org.mindrot" % "jbcrypt" % "0.3m",
  ws,
  "com.typesafe.play.plugins" %% "play-plugins-mailer" % "2.3.0",
  "org.webjars" % "jquery" % "1.11.0",
  "org.webjars" % "angularjs" % "1.2.18",
  "org.webjars" % "bindonce" % "0.3.1",
  "org.webjars" % "ngStorage" % "0.3.0",
  "org.webjars" % "normalize.styl" % "3.0.0",
  "org.webjars" % "resumable.js" % "1.0"
)

resourceGenerators in Compile <+= (resourceManaged in Compile, name, version) map { (dir, name, ver) =>
  val f = dir / "properties"
  val p = new java.util.Properties()
  p.setProperty("name", name)
  p.setProperty("version", ver)
  IO.write(p, "build info", f)
  Seq(f)
}

TwirlKeys.templateImports ++= Seq("macros._", "site._")

JsEngineKeys.engineType := JsEngineKeys.EngineType.Node

pipelineStages in Assets := Seq(uglify)

StylusKeys.useNib in Assets := true

includeFilter in (Assets, StylusKeys.stylus) := "app.styl"

StylusKeys.compress in (Assets, StylusKeys.stylus) := true

JsTaskKeys.jsOptions in (Assets, StylusKeys.stylus) ~= { (s : String) =>
  import spray.json._
  JsObject(s.parseJson.asJsObject.fields
    /* unfortunate hard-coding: */
    .+("paths" -> JsArray(JsString("target/web/web-modules/main/webjars/lib/normalize.styl")))
  ).toString
}

AngularTemplatesKeys.compressRemoveIntertagSpaces := true

AngularTemplatesKeys.naming := { f =>
  new File(f).getName
}

AngularTemplatesKeys.outputHtml := None

includeFilter in uglify := new SimpleFileFilter({ f =>
  f.getPath.startsWith((sourceDirectory in Assets).value.getPath) && f.getName.endsWith(".js")
})

UglifyKeys.uglifyOps := { js =>
  // we assume that app.js is first alphabetically
  Seq((js.sortBy(_._2), "app.min.js"))
}
