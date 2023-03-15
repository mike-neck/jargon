package main

import (
	"os"
	"testing"
)

func TestProperties_UnmarshalXML_AwsSdkPom(t *testing.T) {
	reader, err := os.Open("test/aws-sdk-pom-2.20.24.pom")
	if err != nil {
		t.Fatalf("os.Open test/aws-sdk-pom-2.20.24.pom: %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer reader.Close()
	project, err := ReadPom(reader)
	if err != nil {
		t.Fatalf("ReadPom: %v", err)
	}
	properties := project.Properties
	if properties == nil {
		t.Fatal("project.Properties")
	}
	expected := map[string]string{
		"awsjavasdk.version":                 "${project.version}",
		"awsjavasdk.previous.version":        "2.20.23",
		"jackson.version":                    "2.13.2",
		"jackson.databind.version":           "2.13.4.2",
		"jacksonjr.version":                  "2.13.2",
		"eventstream.version":                "1.0.1",
		"commons.lang.version":               "3.12.0",
		"wiremock.version":                   "2.32.0",
		"slf4j.version":                      "1.7.30",
		"log4j.version":                      "2.17.1",
		"commons.io.version":                 "2.11.0",
		"equalsverifier.version":             "3.7.1",
		"netty.version":                      "4.1.86.Final",
		"unitils.version":                    "3.4.6",
		"xmlunit.version":                    "1.3",
		"project.build.sourceEncoding":       "UTF-8",
		"spotbugs.version":                   "4.2.3",
		"javapoet.verion":                    "1.13.0",
		"org.eclipse.jdt.version":            "3.10.0",
		"org.eclipse.text.version":           "3.5.101",
		"rxjava.version":                     "2.2.21",
		"commons-codec.verion":               "1.15",
		"jmh.version":                        "1.29",
		"awscrt.version":                     "0.21.5",
		"junit5.version":                     "5.8.1",
		"mockito.junit5.version":             "4.6.0",
		"junit4.version":                     "4.13.2",
		"hamcrest.version":                   "1.3",
		"mockito.version":                    "4.3.1",
		"assertj.version":                    "3.20.2",
		"guava.version":                      "29.0-jre",
		"jimfs.version":                      "1.1",
		"testng.version":                     "7.1.0",
		"commons-lang.verson":                "2.6",
		"netty-open-ssl-version":             "2.0.54.Final",
		"dynamodb-local.version":             "1.16.0",
		"sqllite.version":                    "1.0.392",
		"blockhound.version":                 "1.0.6.RELEASE",
		"jetty.version":                      "9.4.45.v20220203",
		"maven.surefire.version":             "3.0.0-M5",
		"maven-compiler-plugin.version":      "3.8.1",
		"maven-checkstyle-plugin.version":    "3.1.2",
		"maven-failsafe-plugin.version":      "3.0.0-M5",
		"maven-jar-plugin.version":           "3.3.0",
		"maven-javadoc-plugin.version":       "3.4.1",
		"maven.build.timestamp.format":       "yyyy",
		"maven-dependency-plugin.version":    "3.1.1",
		"maven-gpg-plugin.version":           "1.6",
		"checkstyle.version":                 "8.42",
		"jacoco-maven-plugin.version":        "0.8.7",
		"nexus-staging-maven-plugin.version": "1.6.8",
		"exec-maven-plugin.version":          "1.6.0",
		"maven-deploy-plugin.version":        "2.8.2",
		"build-helper-maven-plugin.version":  "3.3.0",
		"japicmp-maven-plugin.version":       "0.15.6",
		"versions-maven-plugin.version":      "2.13.0",
		"maven-archetype-plugin.version":     "3.2.1",
		"maven-wrapper-plugin.version":       "3.1.0",
		"json-path.version":                  "2.4.0",
		"spring.version":                     "3.0.7.RELEASE",
		"freemarker.version":                 "2.3.9",
		"aspectj.version":                    "1.8.2",
		"jre.version":                        "1.8.9", // should be failed
		"httpcomponents.httpclient.version":  "4.5.13",
		"httpcomponents.httpcore.version":    "4.4.13",
		"reactive-streams.version":           "1.0.3",
		"skip.unit.tests":                    "${skipTests}",
		"integTestSourceDirectory":           "${project.basedir}/src/it/java",
		"javadoc.resourcesDir":               "${session.executionRootDirectory}",
	}
	for name, value := range expected {
		if properties.Get(name) != value {
			t.Errorf("expected %s for name %s, but got %s", value, name, properties.Get(name))
		}
	}
}
