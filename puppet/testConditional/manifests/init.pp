class lolilol {

  if $facts['os']['family'] == 'Debian' 
  {
    notice("Debian server, we can keep on with the verifications")
    if $facts['os']['distro']['codename'] == 'stretch'
    {
      notice("Stretch release, we can go on...")
    }
  } 


}
