package server

import (
	"log"
	"net/http"

	"github.com/datti-api/pkg/infrastructure/database"
	"github.com/datti-api/pkg/infrastructure/repositoryimpl"
	"github.com/datti-api/pkg/interface/api/handler"
	auth "github.com/datti-api/pkg/interface/api/middleware"
	"github.com/datti-api/pkg/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Sever(dsn string, hostName string, dbInit bool) {
	// DBインスタンスの生成
	dbClient, err := database.NewBunClient(dsn)
	if err != nil {
		log.Print(err.Error())
	}

	tenantClient, err := database.NewFireBaseClient()
	if err != nil {
		log.Print(err.Error())
	}

	// 依存性の解決
	transaction := repositoryimpl.NewTransaction(dbClient.Client)

	userRepository := repositoryimpl.NewProfileRepoImpl(tenantClient)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	bankAccountRepository := repositoryimpl.NewBankAccountRepository(dbClient)
	bankAccountUseCase := usecase.NewBankAccountUseCase(bankAccountRepository, transaction)
	bankAccountHandler := handler.NewBankAccountHandler(bankAccountUseCase)

	friendRepository := repositoryimpl.NewFriendRepository(dbClient)
	friendUseCase := usecase.NewFriendUseCase(friendRepository, userRepository, transaction)
	friendHandler := handler.NewFriendHandler(friendUseCase)

	groupUserRepository := repositoryimpl.NewGroupUserRepository(dbClient)

	groupRepository := repositoryimpl.NewGropuRepoImpl(dbClient)
	groupUseCase := usecase.NewGroupUseCase(groupRepository, userRepository, groupUserRepository, transaction)
	groupHandler := handler.NewGroupHandler(groupUseCase)

	r := echo.New()
	r.Pre(middleware.RemoveTrailingSlash())

	// CORS許可の設定
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodHead},
	}))
	r.Use(auth.FirebaseAuthMiddleware())

	r.GET("/users", userHandler.HandleGetByEmail)
	r.GET("/users/me", userHandler.HandleGetByUid)
	r.PUT("/users/me", userHandler.HandleUpdate)
	r.GET("/users/:uid", userHandler.HandleGetByUidWithPahtParam)
	r.POST("/users/:uid/requests", friendHandler.HandlerRequest) //フレンド申請を送信

	r.GET("/bank", bankAccountHandler.HandleGet)
	r.POST("/bank", bankAccountHandler.HandleUpsert)
	r.DELETE("/bank", bankAccountHandler.HandleDelete)

	r.GET("/friends", friendHandler.HandleGetFriends)            //フレンドを取得
	r.GET("/friends/pendings", friendHandler.HandleGetApplieds)  //フレンド申請未承認のユーザーを取得
	r.GET("/friends/requests", friendHandler.HandleGetApplyings) //フレンド申請中のユーザー
	r.DELETE("/friends/:uid", friendHandler.HandleDelete)        //フレンド登録の解除

	r.GET("/groups", groupHandler.HandleGet)                    //所属グループ一覧の取得
	r.GET("/groups/:id", groupHandler.HandleGetById)            //グループ情報の取得
	r.POST("/groups", groupHandler.HandleCreate)                //グループの作成
	r.PUT("/groups/:id", groupHandler.HandleUpdate)             //グループ情報の更新
	r.POST("/groups/:id/members", groupHandler.HandleRegisterd) //グループに対するメンバーの追加

	r.Start("0.0.0.0:8080")
}
