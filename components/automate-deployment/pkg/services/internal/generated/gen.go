// Code generated by go generate; DO NOT EDIT.
package generated

var ProductMetadataJSON = `
{
  "packages": [
    {
      "name": "chef/backup-gateway",
      "metadata": {
        "name": "chef/backup-gateway",
        "data_service": true,
        "binlinks": null
      }
    },
    {
      "name": "chef/automate-postgresql",
      "metadata": {
        "name": "chef/automate-postgresql",
        "data_service": true,
        "binlinks": null
      }
    },
    {
      "name": "chef/automate-pg-gateway",
      "metadata": null
    },
    {
      "name": "chef/automate-elasticsearch",
      "metadata": {
        "name": "chef/automate-elasticsearch",
        "data_service": true,
        "binlinks": null
      }
    },
    {
      "name": "chef/automate-es-gateway",
      "metadata": null
    },
    {
      "name": "chef/automate-ui",
      "metadata": null
    },
    {
      "name": "chef/pg-sidecar-service",
      "metadata": null
    },
    {
      "name": "chef/event-service",
      "metadata": null
    },
    {
      "name": "chef/authz-service",
      "metadata": null
    },
    {
      "name": "chef/es-sidecar-service",
      "metadata": null
    },
    {
      "name": "chef/automate-dex",
      "metadata": null
    },
    {
      "name": "chef/teams-service",
      "metadata": null
    },
    {
      "name": "chef/authn-service",
      "metadata": null
    },
    {
      "name": "chef/secrets-service",
      "metadata": null
    },
    {
      "name": "chef/applications-service",
      "metadata": null
    },
    {
      "name": "chef/notifications-service",
      "metadata": null
    },
    {
      "name": "chef/nodemanager-service",
      "metadata": null
    },
    {
      "name": "chef/compliance-service",
      "metadata": null
    },
    {
      "name": "chef/license-control-service",
      "metadata": null
    },
    {
      "name": "chef/local-user-service",
      "metadata": null
    },
    {
      "name": "chef/session-service",
      "metadata": null
    },
    {
      "name": "chef/ingest-service",
      "metadata": null
    },
    {
      "name": "chef/config-mgmt-service",
      "metadata": null
    },
    {
      "name": "chef/data-feed-service",
      "metadata": null
    },
    {
      "name": "chef/data-lifecycle-service",
      "metadata": null
    },
    {
      "name": "chef/event-gateway",
      "metadata": null
    },
    {
      "name": "chef/automate-gateway",
      "metadata": null
    },
    {
      "name": "chef/automate-cs-bookshelf",
      "metadata": null
    },
    {
      "name": "chef/automate-cs-oc-bifrost",
      "metadata": null
    },
    {
      "name": "chef/automate-cs-oc-erchef",
      "metadata": null
    },
    {
      "name": "chef/automate-cs-nginx",
      "metadata": {
        "name": "chef/automate-cs-nginx",
        "data_service": false,
        "binlinks": [
          "knife",
          "chef-server-ctl"
        ]
      }
    },
    {
      "name": "chef/automate-workflow-server",
      "metadata": {
        "name": "chef/automate-workflow-server",
        "data_service": false,
        "binlinks": [
          "workflow-ctl"
        ]
      }
    },
    {
      "name": "chef/automate-workflow-nginx",
      "metadata": null
    },
    {
      "name": "chef/automate-load-balancer",
      "metadata": null
    },
    {
      "name": "chef/automate-prometheus",
      "metadata": {
        "name": "chef/automate-prometheus",
        "data_service": true,
        "binlinks": null
      }
    },
    {
      "name": "chef/automate-cli",
      "metadata": {
        "name": "chef/automate-cli",
        "data_service": false,
        "binlinks": [
          "chef-automate"
        ]
      }
    },
    {
      "name": "core/rsync",
      "metadata": null
    }
  ],
  "collections": [
    {
      "name": "core",
      "aliases": null,
      "type": "base",
      "services": [
        "chef/backup-gateway",
        "chef/license-control-service",
        "chef/automate-load-balancer"
      ],
      "packages": [
        "chef/automate-cli",
        "core/rsync"
      ],
      "dependencies": null,
      "hidden": false
    },
    {
      "name": "postgresql",
      "aliases": null,
      "type": "base",
      "services": [
        "chef/automate-postgresql",
        "chef/automate-pg-gateway",
        "chef/pg-sidecar-service"
      ],
      "packages": null,
      "dependencies": null,
      "hidden": false
    },
    {
      "name": "elasticsearch",
      "aliases": null,
      "type": "base",
      "services": [
        "chef/automate-elasticsearch",
        "chef/automate-es-gateway",
        "chef/es-sidecar-service"
      ],
      "packages": null,
      "dependencies": null,
      "hidden": false
    },
    {
      "name": "auth",
      "aliases": null,
      "type": "base",
      "services": [
        "chef/authz-service",
        "chef/authn-service",
        "chef/automate-dex",
        "chef/local-user-service",
        "chef/session-service"
      ],
      "packages": null,
      "dependencies": [
        "core",
        "postgresql"
      ],
      "hidden": false
    },
    {
      "name": "automate",
      "aliases": [
        "automate-full"
      ],
      "type": "product",
      "services": [
        "chef/automate-ui",
        "chef/event-service",
        "chef/teams-service",
        "chef/authn-service",
        "chef/secrets-service",
        "chef/applications-service",
        "chef/notifications-service",
        "chef/nodemanager-service",
        "chef/compliance-service",
        "chef/ingest-service",
        "chef/config-mgmt-service",
        "chef/data-feed-service",
        "chef/data-lifecycle-service",
        "chef/event-gateway",
        "chef/automate-gateway"
      ],
      "packages": null,
      "dependencies": [
        "core",
        "postgresql",
        "elasticsearch",
        "auth"
      ],
      "hidden": false
    },
    {
      "name": "chef-server",
      "aliases": [
        "chef-infra-server"
      ],
      "type": "product",
      "services": [
        "chef/automate-cs-bookshelf",
        "chef/automate-cs-oc-bifrost",
        "chef/automate-cs-oc-erchef",
        "chef/automate-cs-nginx"
      ],
      "packages": null,
      "dependencies": [
        "core",
        "postgresql",
        "elasticsearch"
      ],
      "hidden": false
    },
    {
      "name": "workflow",
      "aliases": null,
      "type": "product",
      "services": [
        "chef/automate-workflow-server",
        "chef/automate-workflow-nginx"
      ],
      "packages": null,
      "dependencies": [
        "automate"
      ],
      "hidden": false
    },
    {
      "name": "monitoring",
      "aliases": null,
      "type": "product",
      "services": [
        "chef/automate-prometheus"
      ],
      "packages": null,
      "dependencies": [
        "automate"
      ],
      "hidden": true
    }
  ]
}
`
