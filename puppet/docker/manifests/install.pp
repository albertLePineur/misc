class docker::install inherits docker {
  package { 'docker-ce':
    ensure => $docker_version
  }
}
