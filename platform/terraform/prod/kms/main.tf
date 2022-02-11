resource "aws_kms_key" "pg" {
  description             = "for main PostgreSQL"
  key_usage               = "ENCRYPT_DECRYPT"
  enable_key_rotation     = true
  is_enabled              = true
  deletion_window_in_days = 30
}

resource "aws_kms_alias" "pg" {
  name          = "alias/enva_pg"
  target_key_id = aws_kms_key.pg.id
}

output "pg_arn" {
  value = aws_kms_key.pg.arn
}
