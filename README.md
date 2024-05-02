# Monitor de Buckets S3

Este é um simples monitor de buckets S3 que exibe informações sobre o último objeto em cada bucket especificado. Utilizo para monitorar backups

## ENV

Crie um arquivo .env na raiz do projeto e defina as seguintes variáveis de ambiente:

AWS_REGION=sua-regiao

AWS_ACCESS_KEY_ID=sua-chave-de-acesso

AWS_SECRET_ACCESS_KEY=sua-chave-secreta

BUCKET_NAMES=lester-backup,mza-backup
