input {
  file {
    path => "/data/*.csv"
    start_position => beginning
    sincedb_path => "/data/.sincedb"
    type => "conviva"
  }
}

filter {

}

output {

  elasticsearch {
        action => "index"
        hosts => "{{.elasticurl}}"
        index => "test-%{+YYYY.MM}"
  }
}
