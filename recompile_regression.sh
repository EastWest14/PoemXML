set -e
cd regression/Java_Regression/src
find . -name '*.class' -delete
javac $(find . -name '*.java' -not -name '*Test.java' -not -type d)
cd ../../..
