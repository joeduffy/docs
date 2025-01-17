---
title: Module ec2transitgateway
---

<div class="section" id="ec2transitgateway">
<h1>ec2transitgateway<a class="headerlink" href="#ec2transitgateway" title="Permalink to this headline">¶</a></h1>
<blockquote>
<div>This provider is a derived work of the <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws">Terraform Provider</a> distributed under
<a class="reference external" href="https://www.mozilla.org/en-US/MPL/2.0/">MPL 2.0</a>. If you encounter a bug or missing feature, first check the
<a class="reference external" href="https://github.com/pulumi/pulumi-aws/issues">pulumi/pulumi-aws repo</a>; however, if that doesn’t turn up
anything, please consult the source <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/issues">terraform-providers/terraform-provider-aws repo</a>.</div></blockquote>
<span class="target" id="module-pulumi_aws.ec2transitgateway"></span><dl class="class">
<dt id="pulumi_aws.ec2transitgateway.AwaitableGetDirectConnectGatewayAttachmentResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">AwaitableGetDirectConnectGatewayAttachmentResult</code><span class="sig-paren">(</span><em>dx_gateway_id=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.AwaitableGetDirectConnectGatewayAttachmentResult" title="Permalink to this definition">¶</a></dt>
<dd></dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.AwaitableGetRouteTableResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">AwaitableGetRouteTableResult</code><span class="sig-paren">(</span><em>default_association_route_table=None</em>, <em>default_propagation_route_table=None</em>, <em>filters=None</em>, <em>id=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.AwaitableGetRouteTableResult" title="Permalink to this definition">¶</a></dt>
<dd></dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.AwaitableGetTransitGatewayResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">AwaitableGetTransitGatewayResult</code><span class="sig-paren">(</span><em>amazon_side_asn=None</em>, <em>arn=None</em>, <em>association_default_route_table_id=None</em>, <em>auto_accept_shared_attachments=None</em>, <em>default_route_table_association=None</em>, <em>default_route_table_propagation=None</em>, <em>description=None</em>, <em>dns_support=None</em>, <em>filters=None</em>, <em>id=None</em>, <em>owner_id=None</em>, <em>propagation_default_route_table_id=None</em>, <em>tags=None</em>, <em>vpn_ecmp_support=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.AwaitableGetTransitGatewayResult" title="Permalink to this definition">¶</a></dt>
<dd></dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.AwaitableGetVpcAttachmentResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">AwaitableGetVpcAttachmentResult</code><span class="sig-paren">(</span><em>dns_support=None</em>, <em>filters=None</em>, <em>id=None</em>, <em>ipv6_support=None</em>, <em>subnet_ids=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>vpc_id=None</em>, <em>vpc_owner_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.AwaitableGetVpcAttachmentResult" title="Permalink to this definition">¶</a></dt>
<dd></dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.AwaitableGetVpnAttachmentResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">AwaitableGetVpnAttachmentResult</code><span class="sig-paren">(</span><em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>vpn_connection_id=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.AwaitableGetVpnAttachmentResult" title="Permalink to this definition">¶</a></dt>
<dd></dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.GetDirectConnectGatewayAttachmentResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">GetDirectConnectGatewayAttachmentResult</code><span class="sig-paren">(</span><em>dx_gateway_id=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetDirectConnectGatewayAttachmentResult" title="Permalink to this definition">¶</a></dt>
<dd><p>A collection of values returned by getDirectConnectGatewayAttachment.</p>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetDirectConnectGatewayAttachmentResult.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetDirectConnectGatewayAttachmentResult.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway Attachment</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetDirectConnectGatewayAttachmentResult.id">
<code class="descname">id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetDirectConnectGatewayAttachmentResult.id" title="Permalink to this definition">¶</a></dt>
<dd><p>id is the provider-assigned unique ID for this managed resource.</p>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.GetRouteTableResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">GetRouteTableResult</code><span class="sig-paren">(</span><em>default_association_route_table=None</em>, <em>default_propagation_route_table=None</em>, <em>filters=None</em>, <em>id=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetRouteTableResult" title="Permalink to this definition">¶</a></dt>
<dd><p>A collection of values returned by getRouteTable.</p>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetRouteTableResult.default_association_route_table">
<code class="descname">default_association_route_table</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetRouteTableResult.default_association_route_table" title="Permalink to this definition">¶</a></dt>
<dd><p>Boolean whether this is the default association route table for the EC2 Transit Gateway</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetRouteTableResult.default_propagation_route_table">
<code class="descname">default_propagation_route_table</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetRouteTableResult.default_propagation_route_table" title="Permalink to this definition">¶</a></dt>
<dd><p>Boolean whether this is the default propagation route table for the EC2 Transit Gateway</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetRouteTableResult.id">
<code class="descname">id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetRouteTableResult.id" title="Permalink to this definition">¶</a></dt>
<dd><p>EC2 Transit Gateway Route Table identifier</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetRouteTableResult.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetRouteTableResult.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway Route Table</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetRouteTableResult.transit_gateway_id">
<code class="descname">transit_gateway_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetRouteTableResult.transit_gateway_id" title="Permalink to this definition">¶</a></dt>
<dd><p>EC2 Transit Gateway identifier</p>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">GetTransitGatewayResult</code><span class="sig-paren">(</span><em>amazon_side_asn=None</em>, <em>arn=None</em>, <em>association_default_route_table_id=None</em>, <em>auto_accept_shared_attachments=None</em>, <em>default_route_table_association=None</em>, <em>default_route_table_propagation=None</em>, <em>description=None</em>, <em>dns_support=None</em>, <em>filters=None</em>, <em>id=None</em>, <em>owner_id=None</em>, <em>propagation_default_route_table_id=None</em>, <em>tags=None</em>, <em>vpn_ecmp_support=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult" title="Permalink to this definition">¶</a></dt>
<dd><p>A collection of values returned by getTransitGateway.</p>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.amazon_side_asn">
<code class="descname">amazon_side_asn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.amazon_side_asn" title="Permalink to this definition">¶</a></dt>
<dd><p>Private Autonomous System Number (ASN) for the Amazon side of a BGP session</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>EC2 Transit Gateway Amazon Resource Name (ARN)</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.association_default_route_table_id">
<code class="descname">association_default_route_table_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.association_default_route_table_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the default association route table</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.auto_accept_shared_attachments">
<code class="descname">auto_accept_shared_attachments</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.auto_accept_shared_attachments" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether resource attachment requests are automatically accepted.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.default_route_table_association">
<code class="descname">default_route_table_association</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.default_route_table_association" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether resource attachments are automatically associated with the default association route table.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.default_route_table_propagation">
<code class="descname">default_route_table_propagation</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.default_route_table_propagation" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether resource attachments automatically propagate routes to the default propagation route table.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.description">
<code class="descname">description</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.description" title="Permalink to this definition">¶</a></dt>
<dd><p>Description of the EC2 Transit Gateway</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.dns_support">
<code class="descname">dns_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.dns_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether DNS support is enabled.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.id">
<code class="descname">id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.id" title="Permalink to this definition">¶</a></dt>
<dd><p>EC2 Transit Gateway identifier</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.owner_id">
<code class="descname">owner_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.owner_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the AWS account that owns the EC2 Transit Gateway</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.propagation_default_route_table_id">
<code class="descname">propagation_default_route_table_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.propagation_default_route_table_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the default propagation route table.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetTransitGatewayResult.vpn_ecmp_support">
<code class="descname">vpn_ecmp_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetTransitGatewayResult.vpn_ecmp_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether VPN Equal Cost Multipath Protocol support is enabled.</p>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">GetVpcAttachmentResult</code><span class="sig-paren">(</span><em>dns_support=None</em>, <em>filters=None</em>, <em>id=None</em>, <em>ipv6_support=None</em>, <em>subnet_ids=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>vpc_id=None</em>, <em>vpc_owner_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult" title="Permalink to this definition">¶</a></dt>
<dd><p>A collection of values returned by getVpcAttachment.</p>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.dns_support">
<code class="descname">dns_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.dns_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether DNS support is enabled.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.id">
<code class="descname">id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.id" title="Permalink to this definition">¶</a></dt>
<dd><p>EC2 Transit Gateway VPC Attachment identifier</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.ipv6_support">
<code class="descname">ipv6_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.ipv6_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether IPv6 support is enabled.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.subnet_ids">
<code class="descname">subnet_ids</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.subnet_ids" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifiers of EC2 Subnets.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway VPC Attachment</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.transit_gateway_id">
<code class="descname">transit_gateway_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.transit_gateway_id" title="Permalink to this definition">¶</a></dt>
<dd><p>EC2 Transit Gateway identifier</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.vpc_id">
<code class="descname">vpc_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.vpc_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 VPC.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.vpc_owner_id">
<code class="descname">vpc_owner_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpcAttachmentResult.vpc_owner_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the AWS account that owns the EC2 VPC.</p>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.GetVpnAttachmentResult">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">GetVpnAttachmentResult</code><span class="sig-paren">(</span><em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>vpn_connection_id=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpnAttachmentResult" title="Permalink to this definition">¶</a></dt>
<dd><p>A collection of values returned by getVpnAttachment.</p>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpnAttachmentResult.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpnAttachmentResult.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway VPN Attachment</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.GetVpnAttachmentResult.id">
<code class="descname">id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.GetVpnAttachmentResult.id" title="Permalink to this definition">¶</a></dt>
<dd><p>id is the provider-assigned unique ID for this managed resource.</p>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.Route">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">Route</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>blackhole=None</em>, <em>destination_cidr_block=None</em>, <em>transit_gateway_attachment_id=None</em>, <em>transit_gateway_route_table_id=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.Route" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages an EC2 Transit Gateway Route.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>blackhole</strong> (<em>pulumi.Input</em><em>[</em><em>bool</em><em>]</em>) – Indicates whether to drop traffic that matches this route (default to <code class="docutils literal notranslate"><span class="pre">false</span></code>).</li>
<li><strong>destination_cidr_block</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.</li>
<li><strong>transit_gateway_attachment_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 Transit Gateway Attachment (required if <code class="docutils literal notranslate"><span class="pre">blackhole</span></code> is set to false).</li>
<li><strong>transit_gateway_route_table_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 Transit Gateway Route Table.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.Route.blackhole">
<code class="descname">blackhole</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.Route.blackhole" title="Permalink to this definition">¶</a></dt>
<dd><p>Indicates whether to drop traffic that matches this route (default to <code class="docutils literal notranslate"><span class="pre">false</span></code>).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.Route.destination_cidr_block">
<code class="descname">destination_cidr_block</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.Route.destination_cidr_block" title="Permalink to this definition">¶</a></dt>
<dd><p>IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.Route.transit_gateway_attachment_id">
<code class="descname">transit_gateway_attachment_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.Route.transit_gateway_attachment_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway Attachment (required if <code class="docutils literal notranslate"><span class="pre">blackhole</span></code> is set to false).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.Route.transit_gateway_route_table_id">
<code class="descname">transit_gateway_route_table_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.Route.transit_gateway_route_table_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway Route Table.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.ec2transitgateway.Route.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>blackhole=None</em>, <em>destination_cidr_block=None</em>, <em>transit_gateway_attachment_id=None</em>, <em>transit_gateway_route_table_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.Route.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Route resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[bool] blackhole: Indicates whether to drop traffic that matches this route (default to <code class="docutils literal notranslate"><span class="pre">false</span></code>).
:param pulumi.Input[str] destination_cidr_block: IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
:param pulumi.Input[str] transit_gateway_attachment_id: Identifier of EC2 Transit Gateway Attachment (required if <code class="docutils literal notranslate"><span class="pre">blackhole</span></code> is set to false).
:param pulumi.Input[str] transit_gateway_route_table_id: Identifier of EC2 Transit Gateway Route Table.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.Route.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.Route.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.Route.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.Route.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.RouteTable">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">RouteTable</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTable" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages an EC2 Transit Gateway Route Table.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Key-value tags for the EC2 Transit Gateway Route Table.</li>
<li><strong>transit_gateway_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 Transit Gateway.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTable.default_association_route_table">
<code class="descname">default_association_route_table</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTable.default_association_route_table" title="Permalink to this definition">¶</a></dt>
<dd><p>Boolean whether this is the default association route table for the EC2 Transit Gateway.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTable.default_propagation_route_table">
<code class="descname">default_propagation_route_table</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTable.default_propagation_route_table" title="Permalink to this definition">¶</a></dt>
<dd><p>Boolean whether this is the default propagation route table for the EC2 Transit Gateway.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTable.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTable.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway Route Table.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTable.transit_gateway_id">
<code class="descname">transit_gateway_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTable.transit_gateway_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.ec2transitgateway.RouteTable.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>default_association_route_table=None</em>, <em>default_propagation_route_table=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTable.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing RouteTable resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[bool] default_association_route_table: Boolean whether this is the default association route table for the EC2 Transit Gateway.
:param pulumi.Input[bool] default_propagation_route_table: Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
:param pulumi.Input[dict] tags: Key-value tags for the EC2 Transit Gateway Route Table.
:param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.RouteTable.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTable.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.RouteTable.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTable.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.RouteTableAssociation">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">RouteTableAssociation</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>transit_gateway_attachment_id=None</em>, <em>transit_gateway_route_table_id=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTableAssociation" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages an EC2 Transit Gateway Route Table association.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>transit_gateway_attachment_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 Transit Gateway Attachment.</li>
<li><strong>transit_gateway_route_table_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 Transit Gateway Route Table.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_association.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_association.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTableAssociation.resource_id">
<code class="descname">resource_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTableAssociation.resource_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the resource</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTableAssociation.resource_type">
<code class="descname">resource_type</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTableAssociation.resource_type" title="Permalink to this definition">¶</a></dt>
<dd><p>Type of the resource</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTableAssociation.transit_gateway_attachment_id">
<code class="descname">transit_gateway_attachment_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTableAssociation.transit_gateway_attachment_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway Attachment.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTableAssociation.transit_gateway_route_table_id">
<code class="descname">transit_gateway_route_table_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTableAssociation.transit_gateway_route_table_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway Route Table.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.ec2transitgateway.RouteTableAssociation.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>resource_id=None</em>, <em>resource_type=None</em>, <em>transit_gateway_attachment_id=None</em>, <em>transit_gateway_route_table_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTableAssociation.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing RouteTableAssociation resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] resource_id: Identifier of the resource
:param pulumi.Input[str] resource_type: Type of the resource
:param pulumi.Input[str] transit_gateway_attachment_id: Identifier of EC2 Transit Gateway Attachment.
:param pulumi.Input[str] transit_gateway_route_table_id: Identifier of EC2 Transit Gateway Route Table.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_association.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_association.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.RouteTableAssociation.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTableAssociation.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.RouteTableAssociation.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTableAssociation.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.RouteTablePropagation">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">RouteTablePropagation</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>transit_gateway_attachment_id=None</em>, <em>transit_gateway_route_table_id=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTablePropagation" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages an EC2 Transit Gateway Route Table propagation.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>transit_gateway_attachment_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 Transit Gateway Attachment.</li>
<li><strong>transit_gateway_route_table_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 Transit Gateway Route Table.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_propagation.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_propagation.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTablePropagation.resource_id">
<code class="descname">resource_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTablePropagation.resource_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the resource</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTablePropagation.resource_type">
<code class="descname">resource_type</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTablePropagation.resource_type" title="Permalink to this definition">¶</a></dt>
<dd><p>Type of the resource</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTablePropagation.transit_gateway_attachment_id">
<code class="descname">transit_gateway_attachment_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTablePropagation.transit_gateway_attachment_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway Attachment.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.RouteTablePropagation.transit_gateway_route_table_id">
<code class="descname">transit_gateway_route_table_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTablePropagation.transit_gateway_route_table_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway Route Table.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.ec2transitgateway.RouteTablePropagation.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>resource_id=None</em>, <em>resource_type=None</em>, <em>transit_gateway_attachment_id=None</em>, <em>transit_gateway_route_table_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTablePropagation.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing RouteTablePropagation resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] resource_id: Identifier of the resource
:param pulumi.Input[str] resource_type: Type of the resource
:param pulumi.Input[str] transit_gateway_attachment_id: Identifier of EC2 Transit Gateway Attachment.
:param pulumi.Input[str] transit_gateway_route_table_id: Identifier of EC2 Transit Gateway Route Table.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_propagation.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_propagation.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.RouteTablePropagation.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTablePropagation.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.RouteTablePropagation.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.RouteTablePropagation.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">TransitGateway</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>amazon_side_asn=None</em>, <em>auto_accept_shared_attachments=None</em>, <em>default_route_table_association=None</em>, <em>default_route_table_propagation=None</em>, <em>description=None</em>, <em>dns_support=None</em>, <em>tags=None</em>, <em>vpn_ecmp_support=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages an EC2 Transit Gateway.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>amazon_side_asn</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is <code class="docutils literal notranslate"><span class="pre">64512</span></code> to <code class="docutils literal notranslate"><span class="pre">65534</span></code> for 16-bit ASNs and <code class="docutils literal notranslate"><span class="pre">4200000000</span></code> to <code class="docutils literal notranslate"><span class="pre">4294967294</span></code> for 32-bit ASNs. Default value: <code class="docutils literal notranslate"><span class="pre">64512</span></code>.</li>
<li><strong>auto_accept_shared_attachments</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Whether resource attachment requests are automatically accepted. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">disable</span></code>.</li>
<li><strong>default_route_table_association</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Whether resource attachments are automatically associated with the default association route table. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</li>
<li><strong>default_route_table_propagation</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</li>
<li><strong>description</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Description of the EC2 Transit Gateway.</li>
<li><strong>dns_support</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Whether DNS support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Key-value tags for the EC2 Transit Gateway.</li>
<li><strong>vpn_ecmp_support</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.amazon_side_asn">
<code class="descname">amazon_side_asn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.amazon_side_asn" title="Permalink to this definition">¶</a></dt>
<dd><p>Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is <code class="docutils literal notranslate"><span class="pre">64512</span></code> to <code class="docutils literal notranslate"><span class="pre">65534</span></code> for 16-bit ASNs and <code class="docutils literal notranslate"><span class="pre">4200000000</span></code> to <code class="docutils literal notranslate"><span class="pre">4294967294</span></code> for 32-bit ASNs. Default value: <code class="docutils literal notranslate"><span class="pre">64512</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.arn">
<code class="descname">arn</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.arn" title="Permalink to this definition">¶</a></dt>
<dd><p>EC2 Transit Gateway Amazon Resource Name (ARN)</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.association_default_route_table_id">
<code class="descname">association_default_route_table_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.association_default_route_table_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the default association route table</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.auto_accept_shared_attachments">
<code class="descname">auto_accept_shared_attachments</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.auto_accept_shared_attachments" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether resource attachment requests are automatically accepted. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">disable</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.default_route_table_association">
<code class="descname">default_route_table_association</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.default_route_table_association" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether resource attachments are automatically associated with the default association route table. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.default_route_table_propagation">
<code class="descname">default_route_table_propagation</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.default_route_table_propagation" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.description">
<code class="descname">description</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.description" title="Permalink to this definition">¶</a></dt>
<dd><p>Description of the EC2 Transit Gateway.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.dns_support">
<code class="descname">dns_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.dns_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether DNS support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.owner_id">
<code class="descname">owner_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.owner_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the AWS account that owns the EC2 Transit Gateway</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.propagation_default_route_table_id">
<code class="descname">propagation_default_route_table_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.propagation_default_route_table_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the default propagation route table</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.vpn_ecmp_support">
<code class="descname">vpn_ecmp_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.vpn_ecmp_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>amazon_side_asn=None</em>, <em>arn=None</em>, <em>association_default_route_table_id=None</em>, <em>auto_accept_shared_attachments=None</em>, <em>default_route_table_association=None</em>, <em>default_route_table_propagation=None</em>, <em>description=None</em>, <em>dns_support=None</em>, <em>owner_id=None</em>, <em>propagation_default_route_table_id=None</em>, <em>tags=None</em>, <em>vpn_ecmp_support=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing TransitGateway resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[float] amazon_side_asn: Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is <code class="docutils literal notranslate"><span class="pre">64512</span></code> to <code class="docutils literal notranslate"><span class="pre">65534</span></code> for 16-bit ASNs and <code class="docutils literal notranslate"><span class="pre">4200000000</span></code> to <code class="docutils literal notranslate"><span class="pre">4294967294</span></code> for 32-bit ASNs. Default value: <code class="docutils literal notranslate"><span class="pre">64512</span></code>.
:param pulumi.Input[str] arn: EC2 Transit Gateway Amazon Resource Name (ARN)
:param pulumi.Input[str] association_default_route_table_id: Identifier of the default association route table
:param pulumi.Input[str] auto_accept_shared_attachments: Whether resource attachment requests are automatically accepted. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">disable</span></code>.
:param pulumi.Input[str] default_route_table_association: Whether resource attachments are automatically associated with the default association route table. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.
:param pulumi.Input[str] default_route_table_propagation: Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.
:param pulumi.Input[str] description: Description of the EC2 Transit Gateway.
:param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.
:param pulumi.Input[str] owner_id: Identifier of the AWS account that owns the EC2 Transit Gateway
:param pulumi.Input[str] propagation_default_route_table_id: Identifier of the default propagation route table
:param pulumi.Input[dict] tags: Key-value tags for the EC2 Transit Gateway.
:param pulumi.Input[str] vpn_ecmp_support: Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.TransitGateway.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.TransitGateway.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">VpcAttachment</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>dns_support=None</em>, <em>ipv6_support=None</em>, <em>subnet_ids=None</em>, <em>tags=None</em>, <em>transit_gateway_default_route_table_association=None</em>, <em>transit_gateway_default_route_table_propagation=None</em>, <em>transit_gateway_id=None</em>, <em>vpc_id=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages an EC2 Transit Gateway VPC Attachment. For examples of custom route table association and propagation, see the EC2 Transit Gateway Networking Examples Guide.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>dns_support</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Whether DNS support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</li>
<li><strong>ipv6_support</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Whether IPv6 support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">disable</span></code>.</li>
<li><strong>subnet_ids</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – Identifiers of EC2 Subnets.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Key-value tags for the EC2 Transit Gateway VPC Attachment.</li>
<li><strong>transit_gateway_default_route_table_association</strong> (<em>pulumi.Input</em><em>[</em><em>bool</em><em>]</em>) – Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.</li>
<li><strong>transit_gateway_default_route_table_propagation</strong> (<em>pulumi.Input</em><em>[</em><em>bool</em><em>]</em>) – Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.</li>
<li><strong>transit_gateway_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 Transit Gateway.</li>
<li><strong>vpc_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – Identifier of EC2 VPC.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.dns_support">
<code class="descname">dns_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.dns_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether DNS support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.ipv6_support">
<code class="descname">ipv6_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.ipv6_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether IPv6 support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">disable</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.subnet_ids">
<code class="descname">subnet_ids</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.subnet_ids" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifiers of EC2 Subnets.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway VPC Attachment.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.transit_gateway_default_route_table_association">
<code class="descname">transit_gateway_default_route_table_association</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.transit_gateway_default_route_table_association" title="Permalink to this definition">¶</a></dt>
<dd><p>Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.transit_gateway_default_route_table_propagation">
<code class="descname">transit_gateway_default_route_table_propagation</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.transit_gateway_default_route_table_propagation" title="Permalink to this definition">¶</a></dt>
<dd><p>Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.transit_gateway_id">
<code class="descname">transit_gateway_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.transit_gateway_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.vpc_id">
<code class="descname">vpc_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.vpc_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 VPC.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.vpc_owner_id">
<code class="descname">vpc_owner_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.vpc_owner_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the AWS account that owns the EC2 VPC.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dns_support=None</em>, <em>ipv6_support=None</em>, <em>subnet_ids=None</em>, <em>tags=None</em>, <em>transit_gateway_default_route_table_association=None</em>, <em>transit_gateway_default_route_table_propagation=None</em>, <em>transit_gateway_id=None</em>, <em>vpc_id=None</em>, <em>vpc_owner_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing VpcAttachment resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">enable</span></code>.
:param pulumi.Input[str] ipv6_support: Whether IPv6 support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>. Default value: <code class="docutils literal notranslate"><span class="pre">disable</span></code>.
:param pulumi.Input[list] subnet_ids: Identifiers of EC2 Subnets.
:param pulumi.Input[dict] tags: Key-value tags for the EC2 Transit Gateway VPC Attachment.
:param pulumi.Input[bool] transit_gateway_default_route_table_association: Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.
:param pulumi.Input[bool] transit_gateway_default_route_table_propagation: Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.
:param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.
:param pulumi.Input[str] vpc_id: Identifier of EC2 VPC.
:param pulumi.Input[str] vpc_owner_id: Identifier of the AWS account that owns the EC2 VPC.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachment.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachment.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter">
<em class="property">class </em><code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">VpcAttachmentAccepter</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>tags=None</em>, <em>transit_gateway_attachment_id=None</em>, <em>transit_gateway_default_route_table_association=None</em>, <em>transit_gateway_default_route_table_propagation=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages the accepter’s side of an EC2 Transit Gateway VPC Attachment.</p>
<p>When a cross-account (requester’s AWS account differs from the accepter’s AWS account) EC2 Transit Gateway VPC Attachment
is created, an EC2 Transit Gateway VPC Attachment resource is automatically created in the accepter’s account.
The requester can use the <code class="docutils literal notranslate"><span class="pre">ec2transitgateway.VpcAttachment</span></code> resource to manage its side of the connection
and the accepter can use the <code class="docutils literal notranslate"><span class="pre">ec2transitgateway.VpcAttachmentAccepter</span></code> resource to “adopt” its side of the
connection into management.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>tags</strong> (<em>pulumi.Input</em><em>[</em><em>dict</em><em>]</em>) – Key-value tags for the EC2 Transit Gateway VPC Attachment.</li>
<li><strong>transit_gateway_attachment_id</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the EC2 Transit Gateway Attachment to manage.</li>
<li><strong>transit_gateway_default_route_table_association</strong> (<em>pulumi.Input</em><em>[</em><em>bool</em><em>]</em>) – Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.</li>
<li><strong>transit_gateway_default_route_table_propagation</strong> (<em>pulumi.Input</em><em>[</em><em>bool</em><em>]</em>) – Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment_accepter.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment_accepter.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.dns_support">
<code class="descname">dns_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.dns_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether DNS support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.ipv6_support">
<code class="descname">ipv6_support</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.ipv6_support" title="Permalink to this definition">¶</a></dt>
<dd><p>Whether IPv6 support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.subnet_ids">
<code class="descname">subnet_ids</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.subnet_ids" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifiers of EC2 Subnets.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.tags">
<code class="descname">tags</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.tags" title="Permalink to this definition">¶</a></dt>
<dd><p>Key-value tags for the EC2 Transit Gateway VPC Attachment.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.transit_gateway_attachment_id">
<code class="descname">transit_gateway_attachment_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.transit_gateway_attachment_id" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the EC2 Transit Gateway Attachment to manage.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.transit_gateway_default_route_table_association">
<code class="descname">transit_gateway_default_route_table_association</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.transit_gateway_default_route_table_association" title="Permalink to this definition">¶</a></dt>
<dd><p>Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.transit_gateway_default_route_table_propagation">
<code class="descname">transit_gateway_default_route_table_propagation</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.transit_gateway_default_route_table_propagation" title="Permalink to this definition">¶</a></dt>
<dd><p>Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.transit_gateway_id">
<code class="descname">transit_gateway_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.transit_gateway_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 Transit Gateway.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.vpc_id">
<code class="descname">vpc_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.vpc_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of EC2 VPC.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.vpc_owner_id">
<code class="descname">vpc_owner_id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.vpc_owner_id" title="Permalink to this definition">¶</a></dt>
<dd><p>Identifier of the AWS account that owns the EC2 VPC.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>dns_support=None</em>, <em>ipv6_support=None</em>, <em>subnet_ids=None</em>, <em>tags=None</em>, <em>transit_gateway_attachment_id=None</em>, <em>transit_gateway_default_route_table_association=None</em>, <em>transit_gateway_default_route_table_propagation=None</em>, <em>transit_gateway_id=None</em>, <em>vpc_id=None</em>, <em>vpc_owner_id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing VpcAttachmentAccepter resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>.
:param pulumi.Input[str] ipv6_support: Whether IPv6 support is enabled. Valid values: <code class="docutils literal notranslate"><span class="pre">disable</span></code>, <code class="docutils literal notranslate"><span class="pre">enable</span></code>.
:param pulumi.Input[list] subnet_ids: Identifiers of EC2 Subnets.
:param pulumi.Input[dict] tags: Key-value tags for the EC2 Transit Gateway VPC Attachment.
:param pulumi.Input[str] transit_gateway_attachment_id: The ID of the EC2 Transit Gateway Attachment to manage.
:param pulumi.Input[bool] transit_gateway_default_route_table_association: Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.
:param pulumi.Input[bool] transit_gateway_default_route_table_propagation: Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: <code class="docutils literal notranslate"><span class="pre">true</span></code>.
:param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.
:param pulumi.Input[str] vpc_id: Identifier of EC2 VPC.
:param pulumi.Input[str] vpc_owner_id: Identifier of the AWS account that owns the EC2 VPC.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment_accepter.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment_accepter.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.translate_output_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of output properties
into a format of their choosing before writing those properties to the resource object.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

<dl class="method">
<dt id="pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.VpcAttachmentAccepter.translate_input_property" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides subclasses of Resource an opportunity to translate names of input properties into
a format of their choosing before sending those properties to the Pulumi engine.</p>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><strong>prop</strong> (<em>str</em>) – A property name.</td>
</tr>
<tr class="field-even field"><th class="field-name">Returns:</th><td class="field-body">A potentially transformed property name.</td>
</tr>
<tr class="field-odd field"><th class="field-name">Return type:</th><td class="field-body">str</td>
</tr>
</tbody>
</table>
</dd></dl>

</dd></dl>

<dl class="function">
<dt id="pulumi_aws.ec2transitgateway.get_direct_connect_gateway_attachment">
<code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">get_direct_connect_gateway_attachment</code><span class="sig-paren">(</span><em>dx_gateway_id=None</em>, <em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>opts=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.get_direct_connect_gateway_attachment" title="Permalink to this definition">¶</a></dt>
<dd><p>Get information on an EC2 Transit Gateway’s attachment to a Direct Connect Gateway.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_dx_gateway_attachment.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_dx_gateway_attachment.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="function">
<dt id="pulumi_aws.ec2transitgateway.get_route_table">
<code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">get_route_table</code><span class="sig-paren">(</span><em>filters=None</em>, <em>id=None</em>, <em>tags=None</em>, <em>opts=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.get_route_table" title="Permalink to this definition">¶</a></dt>
<dd><p>Get information on an EC2 Transit Gateway Route Table.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_route_table.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_route_table.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="function">
<dt id="pulumi_aws.ec2transitgateway.get_transit_gateway">
<code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">get_transit_gateway</code><span class="sig-paren">(</span><em>filters=None</em>, <em>id=None</em>, <em>tags=None</em>, <em>opts=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.get_transit_gateway" title="Permalink to this definition">¶</a></dt>
<dd><p>Get information on an EC2 Transit Gateway.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="function">
<dt id="pulumi_aws.ec2transitgateway.get_vpc_attachment">
<code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">get_vpc_attachment</code><span class="sig-paren">(</span><em>filters=None</em>, <em>id=None</em>, <em>tags=None</em>, <em>opts=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.get_vpc_attachment" title="Permalink to this definition">¶</a></dt>
<dd><p>Get information on an EC2 Transit Gateway VPC Attachment.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_vpc_attachment.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_vpc_attachment.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="function">
<dt id="pulumi_aws.ec2transitgateway.get_vpn_attachment">
<code class="descclassname">pulumi_aws.ec2transitgateway.</code><code class="descname">get_vpn_attachment</code><span class="sig-paren">(</span><em>tags=None</em>, <em>transit_gateway_id=None</em>, <em>vpn_connection_id=None</em>, <em>opts=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_aws.ec2transitgateway.get_vpn_attachment" title="Permalink to this definition">¶</a></dt>
<dd><p>Get information on an EC2 Transit Gateway VPN Attachment.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_vpn_attachment.html.markdown">https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_vpn_attachment.html.markdown</a>.</div></blockquote>
</dd></dl>

</div>
