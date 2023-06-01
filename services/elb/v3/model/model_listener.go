package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 监听器信息。
type Listener struct {

	// 监听器的管理状态。只能设置为true。  不支持该字段，请勿使用。
	AdminStateUp bool `json:"admin_state_up"`

	// 监听器使用的CA证书ID。当且仅当type=client时，才会使用该字段对应的证书。
	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref"`

	// 监听器的最大连接数。  取值：-1表示不限制，默认为-1。  不支持该字段，请勿使用。
	ConnectionLimit int32 `json:"connection_limit"`

	// 监听器的创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如：2021-07-30T12:03:44Z
	CreatedAt string `json:"created_at"`

	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。
	DefaultPoolId string `json:"default_pool_id"`

	// 监听器使用的服务器证书ID。
	DefaultTlsContainerRef string `json:"default_tls_container_ref"`

	// 监听器的描述信息。
	Description string `json:"description"`

	// 客户端与LB之间的HTTPS请求的HTTP2功能的开启状态。 开启后，可提升客户端与LB间的访问性能，但LB与后端服务器间仍采用HTTP1.X协议。  使用说明： - 仅HTTPS协议监听器有效。 - QUIC监听器不能设置该字段，固定返回为true。 - 其他协议的监听器可设置该字段但无效，无论取值如何都不影响监听器正常运行。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	Http2Enable bool `json:"http2_enable"`

	// 监听器ID。
	Id string `json:"id"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers"`

	// 监听器所属的负载均衡器的ID列表。一个监听器只支持关联到一个LB。
	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	// 监听器的名称。  注意：若名称为空，则console上无法选中并查看详情。
	Name string `json:"name"`

	// 监听器所在的项目ID。
	ProjectId string `json:"project_id"`

	// 监听器的监听协议。  [取值：TCP、UDP、HTTP、HTTPS、TERMINATED_HTTPS、QUIC。  使用说明： - 共享型LB上的HTTPS监听器只支持设置为TERMINATED_HTTPS， 创建时传入HTTPS将会自动转为TERMINATED_HTTPS。 - 独享型LB上的HTTPS监听器只支持设置为HTTPS，创建时传入TERMINATED_HTTPS将会自动转为HTTPS。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt)  [取值：TCP、UDP、HTTP、HTTPS。](tag:hws_eu,hcso_dt)  [不支持QUIC。](tag:tm,dt,dt_test)
	Protocol string `json:"protocol"`

	// 监听器的前端监听端口。客户端将请求发送到该端口中。
	ProtocolPort int32 `json:"protocol_port"`

	// 监听器使用的SNI证书（带域名的服务器证书）ID列表。  使用说明： - 列表对应的所有SNI证书的域名不允许存在重复。 - 列表对应的所有SNI证书的域名总数不超过50。
	SniContainerRefs []string `json:"sni_container_refs"`

	// 监听器使用的SNI证书泛域名匹配方式。 longest_suffix表示最长尾缀匹配，wildcard表示标准域名分级匹配。 默认为wildcard。
	SniMatchAlgo string `json:"sni_match_algo"`

	// 标签列表。
	Tags []Tag `json:"tags"`

	// 监听器的更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如：2021-07-30T12:03:44Z
	UpdatedAt string `json:"updated_at"`

	// 监听器使用的安全策略。  [取值：tls-1-0-inherit,tls-1-0, tls-1-1, tls-1-2,tls-1-2-strict，tls-1-2-fs，tls-1-0-with-1-3, tls-1-2-fs-with-1-3, hybrid-policy-1-0，默认：tls-1-0。 ](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,tm,cmcc,dt)  [取值：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict，默认：tls-1-0。](tag:hws_eu,g42,hk_g42,hcso_dt)  [使用说明： - 仅对HTTPS协议类型的监听器且关联LB为独享型时有效。 - QUIC监听器不支持该字段。 - 若同时设置了security_policy_id和tls_ciphers_policy，则仅security_policy_id生效。 - 加密套件的优先顺序为ecc套件、rsa套件、tls1.3协议的套件（即支持ecc又支持rsa） ](tag:hws,hws_hk,hws_eu,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt)  [使用说明： - 仅对HTTPS协议类型的监听器有效](tag:hcso_dt)  [不支持tls1.3协议的套件。](tag:tm,hws_eu,g42,hk_g42)  [不支持QUIC。](tag:tm,dt,dt_test)
	TlsCiphersPolicy string `json:"tls_ciphers_policy"`

	// 自定义安全策略的ID。  [使用说明： - 仅对HTTPS协议类型的监听器且关联LB为独享型时有效。 - 若同时设置了security_policy_id和tls_ciphers_policy，则仅security_policy_id生效。 - 加密套件的优先顺序为ecc套件、rsa套件、tls1.3协议的套件（即支持ecc又支持rsa） ](tag:hws,hws_hk,hws_eu,ocb,ctc,hcso,g42,tm,cmcc,hk-g42,dt)  [使用说明： - 仅对HTTPS协议类型的监听器有效](tag:hcso_dt)  [不支持tls1.3协议的套件。](tag:tm,hws_eu,g42,hk_g42)
	SecurityPolicyId string `json:"security_policy_id"`

	// 是否开启后端服务器的重试。  取值：true 开启重试，false 不开启重试。默认：true。  [使用说明： - 若关联是共享型LB，仅在protocol为HTTP、TERMINATED_HTTPS时才能传入该字段。 - 若关联是独享型LB，仅在protocol为HTTP、HTTPS和QUIC时才能传入该字段。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt)  [使用说明： - 仅在protocol为HTTP、HTTPS时才能传入该字段。](tag:hws_eu,hcso_dt)  [不支持QUIC。](tag:tm,dt,dt_test)
	EnableMemberRetry bool `json:"enable_member_retry"`

	// 客户端连接空闲超时时间。在超过keepalive_timeout时长一直没有请求， 负载均衡会暂时中断当前连接，直到下一次请求时重新建立新的连接。  取值： - 若为TCP监听器，取值范围为（10-4000s）默认值为300s。 - 若为HTTP/HTTPS/TERMINATED_HTTPS监听器，取值范围为（0-4000s）默认值为60s。  UDP监听器不支持此字段。
	KeepaliveTimeout int32 `json:"keepalive_timeout"`

	// 等待客户端请求超时时间，包括两种情况： - 读取整个客户端请求头的超时时长：如果客户端未在超时时长内发送完整个请求头，则请求将被中断 - 两个连续body体的数据包到达LB的时间间隔，超出client_timeout将会断开连接。  取值范围为1-300s，默认值为60s。  使用说明：仅协议为HTTP/HTTPS的监听器支持该字段。
	ClientTimeout int32 `json:"client_timeout"`

	// 等待后端服务器响应超时时间。请求转发后端服务器后，在等待超时member_timeout时长没有响应，负载均衡将终止等待，并返回 HTTP504错误码。  取值：1-300s，默认为60s。  使用说明：仅支持协议为HTTP/HTTPS的监听器。
	MemberTimeout int32 `json:"member_timeout"`

	Ipgroup *ListenerIpGroup `json:"ipgroup"`

	// 是否透传客户端IP地址。开启后客户端IP地址将透传到后端服务器。 [仅作用于共享型LB的TCP/UDP监听器。  取值： - 共享型LB的TCP/UDP监听器可设置为true或false，不传默认为false。 - 共享型LB的HTTP/HTTPS监听器只支持设置为true，不传默认为true。 - 独享型负载均衡器所有协议的监听器只支持设置为true，不传默认为true。  使用说明： - 开启特性后，ELB和后端服务器之间直接使用真实的IP访问，需要确保已正确设置服务器的安全组以及访问控制策略。 - 开启特性后，不支持同一台服务器既作为后端服务器又作为客户端的场景。 - 开启特性后，不支持变更后端服务器规格。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt)  [只设支持置为true，不传默认为true。](tag:hws_eu,hcso_dt)
	TransparentClientIpEnable bool `json:"transparent_client_ip_enable"`

	// 是否开启高级转发策略功能。开启高级转发策略后，支持更灵活的转发策略和转发规则设置。  取值：true开启，false不开启，默认false。  开启后支持如下场景： - 转发策略的action字段支持指定为REDIRECT_TO_URL, FIXED_RESPONSE，即支持URL重定向和响应固定的内容给客户端。 - 转发策略支持指定priority、redirect_url_config、fixed_response_config字段。 - 转发规则rule的type可以指定METHOD, HEADER, QUERY_STRING, SOURCE_IP这几种取值。 - 转发规则rule的type为HOST_NAME时，转发规则rule的value支持通配符*。 - 转发规则支持指定conditions字段。  [荷兰region不支持该字段，请勿使用。](tag:dt)
	EnhanceL7policyEnable bool `json:"enhance_l7policy_enable"`

	QuicConfig *ListenerQuicConfig `json:"quic_config,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *ListenerProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// ELB是否开启gzip压缩，缺省值：false，非必选  [仅HTTP/HTTPS类型监听器支持配置。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test) [仅HTTP/HTTPS/QUIC类型监听器支持配置。](tag:hws,hws_hk,hws_test,hcs,hcs_sm,hcso,fcs,fcs_vm,mix,ocb,ctc,cmcc,sbc,hws_ocb,hk_sbc)
	GzipEnable *bool `json:"gzip_enable,omitempty"`
}

func (o Listener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Listener struct{}"
	}

	return strings.Join([]string{"Listener", string(data)}, " ")
}

type ListenerProtectionStatus struct {
	value string
}

type ListenerProtectionStatusEnum struct {
	NON_PROTECTION     ListenerProtectionStatus
	CONSOLE_PROTECTION ListenerProtectionStatus
}

func GetListenerProtectionStatusEnum() ListenerProtectionStatusEnum {
	return ListenerProtectionStatusEnum{
		NON_PROTECTION: ListenerProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: ListenerProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c ListenerProtectionStatus) Value() string {
	return c.value
}

func (c ListenerProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerProtectionStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
