class docker {
  $docker_version = "5:19.03.4~3-0~debian-stretch"
  $debian_version = "stretch"
  include docker::repo
  include docker::install
}
