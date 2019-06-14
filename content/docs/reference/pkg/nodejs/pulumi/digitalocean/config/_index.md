---
title: Module config
---

<!-- WARNING: this page was generated by a tool. Do not edit it by hand. -->
<!-- To change it, please see https://github.com/pulumi/docs/tree/master/tools/tscdocgen. -->

<a href="../">@pulumi/digitalocean</a> &gt; config

<div class="toggleVisible">
<div class="collapsed">
<h2 class="pdoc-module-header toggleButton" title="Click to show Index">Index ▹</h2>
</div>
<div class="expanded">
<h2 class="pdoc-module-header toggleButton" title="Click to hide Index">Index ▾</h2>
<div class="pdoc-module-contents">
<ul>
<li><a href="#apiEndpoint">let apiEndpoint</a></li>
<li><a href="#spacesAccessId">let spacesAccessId</a></li>
<li><a href="#spacesSecretKey">let spacesSecretKey</a></li>
<li><a href="#token">let token</a></li>
</ul>

<a href="https://github.com/pulumi/pulumi-digitalocean/blob/b6899842c18afdcb0933bffc7c04aa1c74b67e1e/sdk/nodejs/config/vars.ts">config/vars.ts</a> 
</div>
</div>
</div>


<h2 class="pdoc-module-header" id="apiEndpoint">
<a class="pdoc-member-name" href="https://github.com/pulumi/pulumi-digitalocean/blob/b6899842c18afdcb0933bffc7c04aa1c74b67e1e/sdk/nodejs/config/vars.ts#L12">let <b>apiEndpoint</b></a>
</h2>
<div class="pdoc-module-contents">
<pre class="highlight"><span class='kd'>let</span> apiEndpoint: <span class='kd'><a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String'>string</a></span> | <span class='kd'><a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/undefined'>undefined</a></span> = <span class='s2'> __config.get(&#34;apiEndpoint&#34;) || (utilities.getEnv(&#34;DIGITALOCEAN_API_URL&#34;) || &#34;https://api.digitalocean.com&#34;)</span>;</pre>
{{% md %}}

The URL to use for the DigitalOcean API.

{{% /md %}}
</div>
<h2 class="pdoc-module-header" id="spacesAccessId">
<a class="pdoc-member-name" href="https://github.com/pulumi/pulumi-digitalocean/blob/b6899842c18afdcb0933bffc7c04aa1c74b67e1e/sdk/nodejs/config/vars.ts#L16">let <b>spacesAccessId</b></a>
</h2>
<div class="pdoc-module-contents">
<pre class="highlight"><span class='kd'>let</span> spacesAccessId: <span class='kd'><a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String'>string</a></span> | <span class='kd'><a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/undefined'>undefined</a></span> = <span class='s2'> __config.get(&#34;spacesAccessId&#34;) || utilities.getEnv(&#34;SPACES_ACCESS_KEY_ID&#34;)</span>;</pre>
{{% md %}}

The access key ID for Spaces API operations.

{{% /md %}}
</div>
<h2 class="pdoc-module-header" id="spacesSecretKey">
<a class="pdoc-member-name" href="https://github.com/pulumi/pulumi-digitalocean/blob/b6899842c18afdcb0933bffc7c04aa1c74b67e1e/sdk/nodejs/config/vars.ts#L20">let <b>spacesSecretKey</b></a>
</h2>
<div class="pdoc-module-contents">
<pre class="highlight"><span class='kd'>let</span> spacesSecretKey: <span class='kd'><a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String'>string</a></span> | <span class='kd'><a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/undefined'>undefined</a></span> = <span class='s2'> __config.get(&#34;spacesSecretKey&#34;) || utilities.getEnv(&#34;SPACES_SECRET_ACCESS_KEY&#34;)</span>;</pre>
{{% md %}}

The secret access key for Spaces API operations.

{{% /md %}}
</div>
<h2 class="pdoc-module-header" id="token">
<a class="pdoc-member-name" href="https://github.com/pulumi/pulumi-digitalocean/blob/b6899842c18afdcb0933bffc7c04aa1c74b67e1e/sdk/nodejs/config/vars.ts#L24">let <b>token</b></a>
</h2>
<div class="pdoc-module-contents">
<pre class="highlight"><span class='kd'>let</span> token: <span class='kd'><a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String'>string</a></span> | <span class='kd'><a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/undefined'>undefined</a></span> = <span class='s2'> __config.get(&#34;token&#34;) || utilities.getEnv(&#34;DIGITALOCEAN_TOKEN&#34;)</span>;</pre>
{{% md %}}

The token key for API operations.

{{% /md %}}
</div>