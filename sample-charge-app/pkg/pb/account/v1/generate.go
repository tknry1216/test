package accountv1

//go:generate mockgen -destination=../../mock/account/v1/account_service_mock.go -package=mock_accountv1 github.com/flora/pkg/pb/account/v1 AccountServiceServer
