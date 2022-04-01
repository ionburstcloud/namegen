source = ["./namegen"]
bundle_id = "io.ionburst.namegen"

apple_id {
  username = "josh_fraser@ionburst.io"
  password = "@env:AC_PASSWORD"
}

sign {
  application_identity = "1C5E236036FD5489DB56C90EB520A6ECCDAB9363"
  entitlements_file = "./namegen.entitlements"
}

zip {
  output_path = "./namegen.zip"
}
