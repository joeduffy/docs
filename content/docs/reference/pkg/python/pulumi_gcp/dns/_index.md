---
title: Module dns
---

<div class="section" id="dns">
<h1>dns<a class="headerlink" href="#dns" title="Permalink to this headline">¶</a></h1>
<blockquote>
<div>This provider is a derived work of the <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google">Terraform Provider</a> distributed under
<a class="reference external" href="https://www.mozilla.org/en-US/MPL/2.0/">MPL 2.0</a>. If you encounter a bug or missing feature, first check the
<a class="reference external" href="https://github.com/pulumi/pulumi-gcp/issues">pulumi/pulumi-gcp repo</a>; however, if that doesn’t turn up
anything, please consult the source <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/issues">terraform-providers/terraform-provider-google repo</a>.</div></blockquote>
<span class="target" id="module-pulumi_gcp.dns"></span><dl class="class">
<dt id="pulumi_gcp.dns.AwaitableGetManagedZoneResult">
<em class="property">class </em><code class="descclassname">pulumi_gcp.dns.</code><code class="descname">AwaitableGetManagedZoneResult</code><span class="sig-paren">(</span><em>description=None</em>, <em>dns_name=None</em>, <em>name=None</em>, <em>name_servers=None</em>, <em>project=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.AwaitableGetManagedZoneResult" title="Permalink to this definition">¶</a></dt>
<dd></dd></dl>

<dl class="class">
<dt id="pulumi_gcp.dns.GetManagedZoneResult">
<em class="property">class </em><code class="descclassname">pulumi_gcp.dns.</code><code class="descname">GetManagedZoneResult</code><span class="sig-paren">(</span><em>description=None</em>, <em>dns_name=None</em>, <em>name=None</em>, <em>name_servers=None</em>, <em>project=None</em>, <em>id=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.GetManagedZoneResult" title="Permalink to this definition">¶</a></dt>
<dd><p>A collection of values returned by getManagedZone.</p>
<dl class="attribute">
<dt id="pulumi_gcp.dns.GetManagedZoneResult.description">
<code class="descname">description</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.GetManagedZoneResult.description" title="Permalink to this definition">¶</a></dt>
<dd><p>A textual description field.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.dns.GetManagedZoneResult.dns_name">
<code class="descname">dns_name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.GetManagedZoneResult.dns_name" title="Permalink to this definition">¶</a></dt>
<dd><p>The fully qualified DNS name of this zone, e.g. <code class="docutils literal notranslate"><span class="pre">example.com.</span></code>.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.dns.GetManagedZoneResult.name_servers">
<code class="descname">name_servers</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.GetManagedZoneResult.name_servers" title="Permalink to this definition">¶</a></dt>
<dd><p>The list of nameservers that will be authoritative for this
domain. Use NS records to redirect from your DNS provider to these names,
thus making Google Cloud DNS authoritative for this zone.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.dns.GetManagedZoneResult.id">
<code class="descname">id</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.GetManagedZoneResult.id" title="Permalink to this definition">¶</a></dt>
<dd><p>id is the provider-assigned unique ID for this managed resource.</p>
</dd></dl>

</dd></dl>

<dl class="class">
<dt id="pulumi_gcp.dns.ManagedZone">
<em class="property">class </em><code class="descclassname">pulumi_gcp.dns.</code><code class="descname">ManagedZone</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>description=None</em>, <em>dns_name=None</em>, <em>dnssec_config=None</em>, <em>forwarding_config=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>peering_config=None</em>, <em>private_visibility_config=None</em>, <em>project=None</em>, <em>visibility=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.ManagedZone" title="Permalink to this definition">¶</a></dt>
<dd><p>A zone is a subtree of the DNS namespace under one administrative
responsibility. A ManagedZone is a resource that represents a DNS zone
hosted by the Cloud DNS service.</p>
<p>To get more information about ManagedZone, see:</p>
<ul class="simple">
<li><a class="reference external" href="https://cloud.google.com/dns/api/v1/managedZones">API documentation</a></li>
<li>How-to Guides<ul>
<li><a class="reference external" href="https://cloud.google.com/dns/zones/">Managing Zones</a></li>
</ul>
</li>
</ul>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs.
If it is not provided, the provider project is used.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_managed_zone.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_managed_zone.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.dns.ManagedZone.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.ManagedZone.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs.
If it is not provided, the provider project is used.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.dns.ManagedZone.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>description=None</em>, <em>dns_name=None</em>, <em>dnssec_config=None</em>, <em>forwarding_config=None</em>, <em>labels=None</em>, <em>name=None</em>, <em>name_servers=None</em>, <em>peering_config=None</em>, <em>private_visibility_config=None</em>, <em>project=None</em>, <em>visibility=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.ManagedZone.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing ManagedZone resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs.</p>
<blockquote>
<div>If it is not provided, the provider project is used.</div></blockquote>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_managed_zone.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_managed_zone.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.dns.ManagedZone.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.ManagedZone.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.dns.ManagedZone.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.ManagedZone.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.dns.Policy">
<em class="property">class </em><code class="descclassname">pulumi_gcp.dns.</code><code class="descname">Policy</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>alternative_name_server_config=None</em>, <em>description=None</em>, <em>enable_inbound_forwarding=None</em>, <em>enable_logging=None</em>, <em>name=None</em>, <em>networks=None</em>, <em>project=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.Policy" title="Permalink to this definition">¶</a></dt>
<dd><p>A policy is a collection of DNS rules applied to one or more Virtual
Private Cloud resources.</p>
<p>To get more information about Policy, see:</p>
<ul class="simple">
<li><a class="reference external" href="https://cloud.google.com/dns/docs/reference/v1beta2/policies">API documentation</a></li>
<li>How-to Guides<ul>
<li><a class="reference external" href="https://cloud.google.com/dns/zones/#using-dns-server-policies">Using DNS server policies</a></li>
</ul>
</li>
</ul>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs.
If it is not provided, the provider project is used.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_policy.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.dns.Policy.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.Policy.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs.
If it is not provided, the provider project is used.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.dns.Policy.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>alternative_name_server_config=None</em>, <em>description=None</em>, <em>enable_inbound_forwarding=None</em>, <em>enable_logging=None</em>, <em>name=None</em>, <em>networks=None</em>, <em>project=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.Policy.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing Policy resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] project: The ID of the project in which the resource belongs.</p>
<blockquote>
<div>If it is not provided, the provider project is used.</div></blockquote>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_policy.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_policy.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.dns.Policy.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.Policy.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.dns.Policy.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.Policy.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.dns.RecordSet">
<em class="property">class </em><code class="descclassname">pulumi_gcp.dns.</code><code class="descname">RecordSet</code><span class="sig-paren">(</span><em>resource_name</em>, <em>opts=None</em>, <em>managed_zone=None</em>, <em>name=None</em>, <em>project=None</em>, <em>rrdatas=None</em>, <em>ttl=None</em>, <em>type=None</em>, <em>__props__=None</em>, <em>__name__=None</em>, <em>__opts__=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.RecordSet" title="Permalink to this definition">¶</a></dt>
<dd><p>Manages a set of DNS records within Google Cloud DNS. For more information see <a class="reference external" href="https://cloud.google.com/dns/records/">the official documentation</a> and
<a class="reference external" href="https://cloud.google.com/dns/api/v1/resourceRecordSets">API</a>.</p>
<blockquote>
<div><strong>Note:</strong> The provider treats this resource as an authoritative record set. This means existing records (including the default records) for the given type will be overwritten when you create this resource with this provider. In addition, the Google Cloud DNS API requires NS records to be present at all times, so this provider will not actually remove NS records during destroy but will report that it did.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>resource_name</strong> (<em>str</em>) – The name of the resource.</li>
<li><strong>opts</strong> (<a class="reference internal" href="../../pulumi/#pulumi.ResourceOptions" title="pulumi.ResourceOptions"><em>pulumi.ResourceOptions</em></a>) – Options for the resource.</li>
<li><strong>managed_zone</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The name of the zone in which this record set will
reside.</li>
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The DNS name this record set will apply to.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</li>
<li><strong>rrdatas</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – The string data for the records in this record set
whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding <code class="docutils literal notranslate"><span class="pre">&quot;</span></code> if you don’t want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add <code class="docutils literal notranslate"><span class="pre">&quot;&quot;</span></code> inside this provider’s configuration string (e.g. <code class="docutils literal notranslate"><span class="pre">&quot;first255characters&quot;&quot;morecharacters&quot;</span></code>).</li>
<li><strong>ttl</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The time-to-live of this record set (seconds).</li>
<li><strong>type</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The DNS record set type.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_record_set.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_record_set.html.markdown</a>.</div></blockquote>
<dl class="attribute">
<dt id="pulumi_gcp.dns.RecordSet.managed_zone">
<code class="descname">managed_zone</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.managed_zone" title="Permalink to this definition">¶</a></dt>
<dd><p>The name of the zone in which this record set will
reside.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.dns.RecordSet.name">
<code class="descname">name</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.name" title="Permalink to this definition">¶</a></dt>
<dd><p>The DNS name this record set will apply to.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.dns.RecordSet.project">
<code class="descname">project</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.project" title="Permalink to this definition">¶</a></dt>
<dd><p>The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.dns.RecordSet.rrdatas">
<code class="descname">rrdatas</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.rrdatas" title="Permalink to this definition">¶</a></dt>
<dd><p>The string data for the records in this record set
whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding <code class="docutils literal notranslate"><span class="pre">&quot;</span></code> if you don’t want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add <code class="docutils literal notranslate"><span class="pre">&quot;&quot;</span></code> inside this provider’s configuration string (e.g. <code class="docutils literal notranslate"><span class="pre">&quot;first255characters&quot;&quot;morecharacters&quot;</span></code>).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.dns.RecordSet.ttl">
<code class="descname">ttl</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.ttl" title="Permalink to this definition">¶</a></dt>
<dd><p>The time-to-live of this record set (seconds).</p>
</dd></dl>

<dl class="attribute">
<dt id="pulumi_gcp.dns.RecordSet.type">
<code class="descname">type</code><em class="property"> = None</em><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.type" title="Permalink to this definition">¶</a></dt>
<dd><p>The DNS record set type.</p>
</dd></dl>

<dl class="staticmethod">
<dt id="pulumi_gcp.dns.RecordSet.get">
<em class="property">static </em><code class="descname">get</code><span class="sig-paren">(</span><em>resource_name</em>, <em>id</em>, <em>opts=None</em>, <em>managed_zone=None</em>, <em>name=None</em>, <em>project=None</em>, <em>rrdatas=None</em>, <em>ttl=None</em>, <em>type=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.get" title="Permalink to this definition">¶</a></dt>
<dd><p>Get an existing RecordSet resource’s state with the given name, id, and optional extra
properties used to qualify the lookup.
:param str resource_name: The unique name of the resulting resource.
:param str id: The unique provider ID of the resource to lookup.
:param pulumi.ResourceOptions opts: Options for the resource.
:param pulumi.Input[str] managed_zone: The name of the zone in which this record set will</p>
<blockquote>
<div>reside.</div></blockquote>
<table class="docutils field-list" frame="void" rules="none">
<col class="field-name" />
<col class="field-body" />
<tbody valign="top">
<tr class="field-odd field"><th class="field-name">Parameters:</th><td class="field-body"><ul class="first last simple">
<li><strong>name</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The DNS name this record set will apply to.</li>
<li><strong>project</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The ID of the project in which the resource belongs. If it
is not provided, the provider project is used.</li>
<li><strong>rrdatas</strong> (<em>pulumi.Input</em><em>[</em><em>list</em><em>]</em>) – The string data for the records in this record set
whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding <code class="docutils literal notranslate"><span class="pre">&quot;</span></code> if you don’t want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add <code class="docutils literal notranslate"><span class="pre">&quot;&quot;</span></code> inside this provider’s configuration string (e.g. <code class="docutils literal notranslate"><span class="pre">&quot;first255characters&quot;&quot;morecharacters&quot;</span></code>).</li>
<li><strong>ttl</strong> (<em>pulumi.Input</em><em>[</em><em>float</em><em>]</em>) – The time-to-live of this record set (seconds).</li>
<li><strong>type</strong> (<em>pulumi.Input</em><em>[</em><em>str</em><em>]</em>) – The DNS record set type.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_record_set.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_record_set.html.markdown</a>.</div></blockquote>
</dd></dl>

<dl class="method">
<dt id="pulumi_gcp.dns.RecordSet.translate_output_property">
<code class="descname">translate_output_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.translate_output_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.dns.RecordSet.translate_input_property">
<code class="descname">translate_input_property</code><span class="sig-paren">(</span><em>prop</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.RecordSet.translate_input_property" title="Permalink to this definition">¶</a></dt>
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
<dt id="pulumi_gcp.dns.get_managed_zone">
<code class="descclassname">pulumi_gcp.dns.</code><code class="descname">get_managed_zone</code><span class="sig-paren">(</span><em>name=None</em>, <em>project=None</em>, <em>opts=None</em><span class="sig-paren">)</span><a class="headerlink" href="#pulumi_gcp.dns.get_managed_zone" title="Permalink to this definition">¶</a></dt>
<dd><p>Provides access to a zone’s attributes within Google Cloud DNS.
For more information see
<a class="reference external" href="https://cloud.google.com/dns/zones/">the official documentation</a>
and
<a class="reference external" href="https://cloud.google.com/dns/api/v1/managedZones">API</a>.</p>
<blockquote>
<div>This content is derived from <a class="reference external" href="https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/dns_managed_zone.html.markdown">https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/dns_managed_zone.html.markdown</a>.</div></blockquote>
</dd></dl>

</div>
