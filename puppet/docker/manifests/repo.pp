class docker::repo inherits docker {

  include apt

  apt::source { 'docker-ce':
    comment  => 'Docker Community Edition for Debian',
    location => 'https://download.docker.com/linux/debian',
    release  => $debian_version,
    repos    => 'stable',
    architecture => 'amd64',
    key      => {
      'id'     => '9DC858229FC7DD38854AE2D88D81803C0EBFCD88',
      'source' => 'https://download.docker.com/linux/debian/gpg',
    },
    include  => {
      'src' => false,
      'deb' => true,
    },
  }
}
