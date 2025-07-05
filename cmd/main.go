package main

import (
	"log"

	"github.com/dedimurphy/blog-api/internal/configs"
	"github.com/dedimurphy/blog-api/internal/handlers/memberships"
  "github.com/dedimurphy/blog-api/internal/handlers/posts"
	membershipsRepo "github.com/dedimurphy/blog-api/internal/repository/memberships"
	membershipsSvc "github.com/dedimurphy/blog-api/internal/service/memberships"
  postSvc "github.com/dedimurphy/blog-api/internal/service/posts"
  postRepo "github.com/dedimurphy/blog-api/internal/repository/posts"
	"github.com/dedimurphy/blog-api/pkg/internalsql"
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

