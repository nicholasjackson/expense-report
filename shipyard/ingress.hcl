ingress "mysql" {
  source {
    driver = "local"

    config {
      port = 3306
    }
  }

  destination {
    driver = "k8s"

    config {
      cluster = "k8s_cluster.dc1"
      address = "expense-db-mysql.default.svc"
      port    = 3306
    }
  }
}

ingress "expense" {
  source {
    driver = "local"

    config {
      port = 5001
    }
  }

  destination {
    driver = "k8s"

    config {
      cluster = "k8s_cluster.dc1"
      address = "expense.default.svc"
      port    = 5001
    }
  }
}

ingress "report" {
  source {
    driver = "local"

    config {
      port = 15001
    }
  }

  destination {
    driver = "k8s"

    config {
      cluster = "k8s_cluster.dc1"
      address = "report.default.svc"
      port    = 5001
    }
  }
}