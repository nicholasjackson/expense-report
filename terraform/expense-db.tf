resource "kubernetes_service_account" "expense_db_mysql" {
  metadata {
    name = "expense-db-mysql"
  }

  automount_service_account_token = true
}

resource "kubernetes_deployment" "expense_db_mysql" {
  metadata {
    name = "expense-db-mysql"
    labels = {
      app     = "expense-db-mysql"
      release = "v1"
    }
  }

  spec {
    replicas = 3

    selector {
      match_labels = {
        app     = "expense-db-mysql"
        release = "v1"
      }
    }

    template {
      metadata {
        labels = {
          app     = "expense-db-mysql"
          release = "v1"
        }
      }

      spec {
        container {
          image                = "joatmon08/expense-db:mysql-8"
          name                 = "expense-db-mysql"
          service_account_name = kubernetes_service_account.expense_db_mysql.metadata.0.name

          port {
            container_port = 3306
          }

          env {
            name  = "MYSQL_ROOT_PASSWORD"
            value = "password"
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }

          liveness_probe {
            tcp_socket {
              port = 3306
            }

            initial_delay_seconds = 3
            period_seconds        = 3
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "expense_db_mysql" {
  metadata {
    name = "expense-db-mysql"
  }
  spec {
    selector = {
      app = kubernetes_deployment.expense_db_mysql.metadata.0.labels.app
    }
    port {
      port        = 3306
      target_port = 3306
    }
  }
}