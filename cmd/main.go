package main

import (
	"log"

	"github.com/dedimurphy/fast-campus/internal/configs"
	"github.com/dedimurphy/fast-campus/internal/handlers/memberships"
  "github.com/dedimurphy/fast-campus/internal/handlers/posts"
	membershipsRepo "github.com/dedimurphy/fast-campus/internal/repository/memberships"
	membershipsSvc "github.com/dedimurphy/fast-campus/internal/service/memberships"
  postSvc "github.com/dedimurphy/fast-campus/internal/service/posts"
  postRepo "github.com/dedimurphy/fast-campus/internal/repository/posts"
	"github.com/dedimurphy/fast-campus/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  var (
    cfg *configs.Config
  )

  err:= configs.Init(
    configs.WithConfigFolder(
      []string{"./internal/configs"},
    ),
    configs.WithConfigFile("config"),
    configs.WithConfigType("yaml"),
  )

  if err != nil {
    log.Fatal("Gagal melakukan inisiasi")
  }

  cfg = configs.Get()
  log.Println("config", cfg)

  db, err := internalsql.Connect(cfg.Database.DataSourceName)
  if err != nil {
    log.Fatal("Gagal inisiasi Database", err)
  }

  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  membershipsRepo := membershipsRepo.NewRepository(db)
  postRepo :=  postRepo.NewRepository(db)

  membershipService := membershipsSvc.NewService(cfg, membershipsRepo)
  postService := postSvc.NewService(cfg, postRepo)

  membershipsHandler := memberships.NewHandler(r, membershipService)
  postHandler := posts.NewHandler(r, postService)

  membershipsHandler.RegisterRoute()
  postHandler.RegisterRoute()
 
  r.Run(cfg.Service.Port) 
}

