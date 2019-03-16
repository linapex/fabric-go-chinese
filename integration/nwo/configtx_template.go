
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456078112854016>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package nwo

const DefaultConfigTxTemplate = `---
{{ with $w := . -}}
Organizations:{{ range .PeerOrgs }}
- &{{ .MSPID }}
  Name: {{ .Name }}
  ID: {{ .MSPID }}
  MSPDir: {{ $w.PeerOrgMSPDir . }}
  Policies:
    Readers:
      Type: Signature
      Rule: OR('{{.MSPID}}.admin', '{{.MSPID}}.peer', '{{.MSPID}}.client')
    Writers:
      Type: Signature
      Rule: OR('{{.MSPID}}.admin', '{{.MSPID}}.client')
    Admins:
      Type: Signature
      Rule: OR('{{.MSPID}}.admin')
  AnchorPeers:{{ range $w.AnchorsInOrg .Name }}
  - Host: 127.0.0.1
    Port: {{ $w.PeerPort . "Listen" }}
  {{- end }}
{{- end }}
{{- range .OrdererOrgs }}
- &{{ .MSPID }}
  Name: {{ .Name }}
  ID: {{ .MSPID }}
  MSPDir: {{ $w.OrdererOrgMSPDir . }}
  Policies:
    Readers:
      Type: Signature
      Rule: OR('{{.MSPID}}.member')
    Writers:
      Type: Signature
      Rule: OR('{{.MSPID}}.member')
    Admins:
      Type: Signature
      Rule: OR('{{.MSPID}}.admin')
{{ end }}

Channel: &ChannelDefaults
  Capabilities:
    V1_3: true
  Policies:
    Readers:
      Type: ImplicitMeta
      Rule: ANY Readers
    Writers:
      Type: ImplicitMeta
      Rule: ANY Writers
    Admins:
      Type: ImplicitMeta
      Rule: MAJORITY Admins

Profiles:{{ range .Profiles }}
  {{ .Name }}:
    {{- if .Orderers }}
    <<: *ChannelDefaults
    Consortiums:{{ range $w.Consortiums }}
      {{ .Name }}:
        Organizations:{{ range .Organizations }}
        - *{{ ($w.Organization .).MSPID }}
        {{- end }}
    {{- end }}
    Orderer:
      OrdererType: {{ $w.Consensus.Type }}
      Addresses:{{ range .Orderers }}{{ with $w.Orderer . }}
      - 127.0.0.1:{{ $w.OrdererPort . "Listen" }}
      {{- end }}{{ end }}
      BatchTimeout: 1s
      BatchSize:
        MaxMessageCount: 1
        AbsoluteMaxBytes: 98 MB
        PreferredMaxBytes: 512 KB
      Capabilities:
        V1_1: true
      {{- if eq $w.Consensus.Type "kafka" }}
      Kafka:
        Brokers:{{ range $w.BrokerAddresses "HostPort" }}
        - {{ . }}
        {{- end }}
      {{- end }}
      {{- if eq $w.Consensus.Type "etcdraft" }}
      EtcdRaft:
        Options:
          SnapshotInterval: 5
        Consenters:{{ range .Orderers }}{{ with $w.Orderer . }}
        - Host: 127.0.0.1
          Port: {{ $w.OrdererPort . "Listen" }}
          ClientTLSCert: {{ $w.OrdererLocalCryptoDir . "tls" }}/server.crt
          ServerTLSCert: {{ $w.OrdererLocalCryptoDir . "tls" }}/server.crt
        {{- end }}{{- end }}
      {{- end }}
      Organizations:{{ range $w.OrgsForOrderers .Orderers }}
      - *{{ .MSPID }}
      {{- end }}
      Policies:
        Readers:
          Type: ImplicitMeta
          Rule: ANY Readers
        Writers:
          Type: ImplicitMeta
          Rule: ANY Writers
        Admins:
          Type: ImplicitMeta
          Rule: MAJORITY Admins
        BlockValidation:
          Type: ImplicitMeta
          Rule: ANY Writers
    {{- else }}
    Application:
      Capabilities:
        V1_3: true
        CAPABILITY_PLACEHOLDER: false
      Organizations:{{ range .Organizations }}
      - *{{ ($w.Organization .).MSPID }}
      {{- end}}
      Policies:
        Readers:
          Type: ImplicitMeta
          Rule: ANY Readers
        Writers:
          Type: ImplicitMeta
          Rule: ANY Writers
        Admins:
          Type: ImplicitMeta
          Rule: MAJORITY Admins
    Consortium: {{ .Consortium }}
    {{- end }}
{{- end }}
{{ end }}
`

