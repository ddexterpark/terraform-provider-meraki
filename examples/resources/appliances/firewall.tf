locals {
  l3firewall-rules : [
    {
      "comment": "Allow TCP traffic to subnet with HTTPS servers.",
      "policy": "allow",
      "protocol": "tcp",
      "destPort": "443",
      "destCidr": "192.168.1.0/24",
      "srcPort": "Any",
      "srcCidr": "Any",
      "syslogEnabled": false
    },
    {
      "comment": "Deny TCP traffic to subnet with HTTP servers.",
      "policy": "deny",
      "protocol": "tcp",
      "destPort": "80",
      "destCidr": "192.168.1.0/24",
      "srcPort": "Any",
      "srcCidr": "Any",
      "syslogEnabled": true
    }
  ]
}

resource "meraki_appliance_l3firewallrules" "l3rules" {
  networkId = ""
  deviceId = ""

  dynamic "rules" {
    for_each = local.l3firewall-rules
    content {
      comment = rules.value.comment
      policy = rules.value.policy
      protocol = rules.value.protocol
      destPort = rules.value.destPort
      destCidr = rules.value.destCidr
      srcPort = rules.value.srcPort
      srcCidr = rules.value.srcCidr
      syslogEnabled = rules.value.syslogEnabled
    }
  }
}