package service

import (
	"log"

	"watch/backend/internal/models"

	"gorm.io/gorm"
)

func SeedContentPages(db *gorm.DB) {
	var count int64
	db.Model(&models.ContentPage{}).Count(&count)
	if count > 0 {
		return
	}

	log.Println("seeding content pages...")

	pages := []models.ContentPage{
		{Slug: "shopping-flow", Title: "购物流程", Category: "购物指南", Sort: 1, Enabled: true, Body: shoppingFlowBody},
		{Slug: "payment-methods", Title: "支付方式", Category: "购物指南", Sort: 2, Enabled: true, Body: paymentMethodsBody},
		{Slug: "order-query", Title: "订单查询", Category: "购物指南", Sort: 3, Enabled: true, Body: orderQueryBody},
		{Slug: "delivery", Title: "配送说明", Category: "购物指南", Sort: 4, Enabled: true, Body: deliveryBody},
		{Slug: "return-policy", Title: "退换政策", Category: "售后服务", Sort: 1, Enabled: true, Body: returnPolicyBody},
		{Slug: "warranty", Title: "保修条款", Category: "售后服务", Sort: 2, Enabled: true, Body: warrantyBody},
		{Slug: "service-centers", Title: "维修网点", Category: "售后服务", Sort: 3, Enabled: true, Body: serviceCentersBody},
		{Slug: "faq", Title: "常见问题", Category: "售后服务", Sort: 4, Enabled: true, Body: faqBody},
		{Slug: "brand-story", Title: "品牌故事", Category: "关于我们", Sort: 1, Enabled: true, Body: brandStoryBody},
		{Slug: "stores", Title: "门店地址", Category: "关于我们", Sort: 2, Enabled: true, Body: storesBody},
		{Slug: "business", Title: "商务合作", Category: "关于我们", Sort: 3, Enabled: true, Body: businessBody},
		{Slug: "contact", Title: "联系我们", Category: "关于我们", Sort: 4, Enabled: true, Body: contactBody},
		{Slug: "privacy", Title: "隐私政策", Category: "法律条款", Sort: 1, Enabled: true, Body: privacyBody},
		{Slug: "terms", Title: "用户协议", Category: "法律条款", Sort: 2, Enabled: true, Body: termsBody},
	}

	for _, p := range pages {
		db.Create(&p)
	}
	log.Println("content pages seed completed")
}

const shoppingFlowBody = `<h2>购物流程</h2>
<p>在胖子腕表选购心仪腕表，只需简单四步：</p>
<ol>
<li><strong>浏览选购</strong> — 在首页按品牌筛选，或进入商品详情页查看规格、图片与用户评价。</li>
<li><strong>确认款式</strong> — 选择表盘颜色、表带款式等规格，确认价格后点击「立即购买」。</li>
<li><strong>填写订单</strong> — 填写收货人姓名、联系电话、收货地址及备注信息，核对无误后提交订单。</li>
<li><strong>完成支付</strong> — 选择微信、支付宝或银行卡支付，支付成功后我们将尽快为您安排发货。</li>
</ol>
<p>如需代购其他品牌款式，请选择「联系客服下单」，专属顾问将一对一为您服务。</p>`

const paymentMethodsBody = `<h2>支付方式</h2>
<p>我们支持以下安全便捷的支付方式：</p>
<ul>
<li><strong>微信支付</strong> — 扫码即可完成支付，支持信用卡绑定。</li>
<li><strong>支付宝</strong> — 支持余额、花呗及绑定的银行卡。</li>
<li><strong>银行卡支付</strong> — 支持主流借记卡与信用卡，由银联提供安全保障。</li>
</ul>
<h3>支付说明</h3>
<ul>
<li>订单提交后请在 <strong>30 分钟</strong> 内完成支付，超时订单将自动取消。</li>
<li>支付成功后，您将收到订单确认信息，可在页脚「订单查询」处跟踪物流状态。</li>
<li>所有交易均通过加密通道传输，我们不会存储您的完整银行卡信息。</li>
</ul>`

const orderQueryBody = `<h2>订单查询</h2>
<p>您可以通过以下方式查询订单状态：</p>
<h3>在线查询</h3>
<p>在网站页脚「订单查询」区域，输入下单时填写的<strong>邮箱</strong>和<strong>预留手机号</strong>，即可查看订单进度。</p>
<h3>客服查询</h3>
<p>拨打客服热线 <strong>400-888-6688</strong>（9:00–21:00），提供订单号或手机号，客服人员将为您查询。</p>
<h3>订单状态说明</h3>
<ul>
<li><strong>待支付</strong> — 订单已创建，等待付款</li>
<li><strong>已支付</strong> — 付款成功，仓库备货中</li>
<li><strong>已发货</strong> — 商品已发出，可查看物流单号</li>
<li><strong>已完成</strong> — 您已确认收货</li>
</ul>`

const deliveryBody = `<h2>配送说明</h2>
<h3>配送范围</h3>
<p>全国范围内（港澳台及偏远地区除外）均可配送。部分偏远乡镇可能需延长 1–2 个工作日。</p>
<h3>配送时效</h3>
<ul>
<li>现货商品：付款后 <strong>1–3 个工作日</strong> 发出</li>
<li>预订/调货商品：以商品页标注或客服通知为准，通常 7–15 个工作日</li>
<li>顺丰速运包邮（订单满 ¥5,000 免运费，不足部分收取 ¥15 运费）</li>
</ul>
<h3>签收须知</h3>
<ul>
<li>签收前请当面验货，确认外包装完好、配件齐全。</li>
<li>高价值商品建议本人签收，必要时可要求快递员配合开箱验货。</li>
<li>如外包装有明显破损，请拍照留证并联系客服，勿直接签收。</li>
</ul>`

const returnPolicyBody = `<h2>退换政策</h2>
<h3>7 天无理由退货</h3>
<p>自签收之日起 <strong>7 日内</strong>，商品未佩戴、未拆封保护膜、配件及包装完好，可申请无理由退货。定制刻字、特殊订货款除外。</p>
<h3>15 天换货</h3>
<p>签收后 <strong>15 日内</strong>，如商品存在非人为质量问题，可申请换货。经鉴定确认后，我们将为您更换同款或等价商品。</p>
<h3>不适用退换的情况</h3>
<ul>
<li>已佩戴、表带调整或存在使用痕迹</li>
<li>人为损坏、进水、私自拆卸</li>
<li>缺少原装配件、保卡或防伪标签</li>
<li>超过退换期限</li>
</ul>
<h3>退换流程</h3>
<p>联系客服（400-888-6688）→ 获取退货地址 → 寄回商品 → 质检通过后 3–7 个工作日退款至原支付账户。</p>`

const warrantyBody = `<h2>保修条款</h2>
<h3>保修范围</h3>
<p>胖子腕表所售腕表均享受<strong>国际联保 2 年</strong>（以品牌官方政策为准）。保修涵盖机芯、走时不准等非人为故障。</p>
<h3>保修凭证</h3>
<p>请妥善保管<strong>保卡、购买凭证</strong>及<strong>防伪标签</strong>，送修时一并提供。</p>
<h3>不在保修范围内</h3>
<ul>
<li>正常磨损（表带、表冠、表镜划痕等）</li>
<li>人为撞击、进水、磁化、私自拆修</li>
<li>不可抗力造成的损坏</li>
<li>超出保修期限</li>
</ul>
<h3>延保服务</h3>
<p>部分品牌支持购买官方延保，详情咨询客服或到店了解。</p>`

const serviceCentersBody = `<h2>维修网点</h2>
<p>胖子腕表与各大品牌授权服务中心合作，为您提供专业维修与保养服务。</p>
<h3>北京服务中心</h3>
<p>地址：北京市朝阳区建国门外大街 1 号国贸商城 B1 层<br>电话：010-8516-xxxx<br>营业时间：10:00–20:00（周一至周日）</p>
<h3>上海服务中心</h3>
<p>地址：上海市静安区南京西路 1266 号恒隆广场 2 层<br>电话：021-6288-xxxx<br>营业时间：10:00–21:00（周一至周日）</p>
<h3>广州服务中心</h3>
<p>地址：广州市天河区天河路 208 号天河城 3 层<br>电话：020-3888-xxxx<br>营业时间：10:00–22:00（周一至周日）</p>
<p>其他城市用户可将腕表寄至以上服务中心，或拨打 <strong>400-888-6688</strong> 查询就近授权网点。</p>`

const faqBody = `<h2>常见问题</h2>
<h3>Q：你们出售的是正品吗？</h3>
<p>A：是的。胖子腕表所有商品均为品牌授权渠道采购，附原装保卡与防伪标识，支持专柜验货。</p>
<h3>Q：可以开发票吗？</h3>
<p>A：可以。下单时在备注中说明发票抬头及税号，我们将随货寄出电子或纸质发票。</p>
<h3>Q：腕表每天慢/快几秒正常吗？</h3>
<p>A：机械表日差 ±15 秒以内属于正常范围。如超出较多，建议到授权服务中心检测。</p>
<h3>Q：如何保养机械表？</h3>
<p>A：避免剧烈碰撞和磁场；建议 3–5 年进行一次洗油保养；不佩戴时建议放入表盒或摇表器。</p>
<h3>Q：支持以旧换新吗？</h3>
<p>A：部分活动期支持，请关注官网公告或咨询客服了解当前政策。</p>`

const brandStoryBody = `<h2>品牌故事</h2>
<p>胖子腕表创立于 2018 年，源于创始人对机械腕表的热爱与对正品保障的坚持。我们相信，每一枚腕表都承载着时间的重量与佩戴者的故事。</p>
<p>从最初的一家社区表店，到如今服务全国数万表友，胖子腕表始终坚持<strong>只卖正品、透明定价、专业售后</strong>三大原则。我们与劳力士、欧米茄、卡地亚等主流品牌保持授权合作，让每位顾客都能放心选购。</p>
<p>「让每一秒都值得珍藏」—— 这不只是我们的 slogan，更是我们对每一位顾客的承诺。</p>`

const storesBody = `<h2>门店地址</h2>
<h3>北京旗舰店</h3>
<p>北京市朝阳区建国门外大街 1 号国贸商城 B1 层<br>电话：010-8516-xxxx<br>营业时间：10:00–20:00</p>
<h3>上海体验店</h3>
<p>上海市静安区南京西路 1266 号恒隆广场 2 层<br>电话：021-6288-xxxx<br>营业时间：10:00–21:00</p>
<h3>深圳店</h3>
<p>深圳市福田区福华三路 269 号 COCOPark B1 层<br>电话：0755-8888-xxxx<br>营业时间：10:00–22:00</p>
<p>欢迎到店试戴体验，门店顾问将为您提供一对一选表建议。</p>`

const businessBody = `<h2>商务合作</h2>
<p>胖子腕表欢迎以下类型的商务合作：</p>
<ul>
<li><strong>品牌授权与经销</strong> — 正规渠道品牌方、代理商</li>
<li><strong>企业团购</strong> — 公司年会、客户礼品、员工福利采购</li>
<li><strong>媒体与 KOL</strong> — 腕表评测、内容合作、联名活动</li>
<li><strong>异业联盟</strong> — 高端生活方式品牌跨界合作</li>
</ul>
<h3>联系方式</h3>
<p>商务邮箱：business@pangziwatch.com<br>电话：400-888-6688 转 2<br>我们将在 3 个工作日内回复您的合作意向。</p>`

const contactBody = `<h2>联系我们</h2>
<h3>客服热线</h3>
<p><strong>400-888-6688</strong><br>服务时间：每日 9:00 – 21:00</p>
<h3>在线客服</h3>
<p>网站右下角在线客服（工作时间同热线）</p>
<h3>电子邮箱</h3>
<p>售前咨询：service@pangziwatch.com<br>售后投诉：support@pangziwatch.com</p>
<h3>公司地址</h3>
<p>北京市朝阳区建国门外大街 1 号国贸写字楼 A 座 18 层<br>邮编：100020</p>`

const privacyBody = `<h2>隐私政策</h2>
<p>胖子腕表（以下简称「我们」）重视您的隐私。本政策说明我们如何收集、使用和保护您的个人信息。</p>
<h3>信息收集</h3>
<p>我们可能收集：姓名、电话、邮箱、收货地址、订单信息及浏览记录，用于完成交易与改善服务。</p>
<h3>信息使用</h3>
<p>您的信息仅用于订单处理、物流配送、售后服务及经您同意的营销推广，不会出售给第三方。</p>
<h3>信息安全</h3>
<p>我们采用 SSL 加密传输及访问控制措施保护您的数据安全。</p>
<h3>您的权利</h3>
<p>您有权查询、更正或删除个人信息，请联系 service@pangziwatch.com。</p>
<p>本政策最后更新：2026 年 1 月</p>`

const termsBody = `<h2>用户协议</h2>
<p>欢迎使用胖子腕表网站。访问或使用本网站即表示您同意以下条款：</p>
<h3>服务说明</h3>
<p>本网站提供腕表商品展示、在线下单及信息查询服务。商品信息以页面描述为准，我们保留调整价格与库存的权利。</p>
<h3>用户义务</h3>
<ul>
<li>提供真实、准确的注册与收货信息</li>
<li>不得利用本网站从事违法活动</li>
<li>尊重知识产权，不得擅自复制网站内容</li>
</ul>
<h3>免责声明</h3>
<p>因不可抗力、网络故障等导致的服务中断，我们将尽力恢复但不承担由此产生的间接损失。</p>
<h3>争议解决</h3>
<p>本协议适用中华人民共和国法律。如有争议，双方应友好协商；协商不成，提交北京市朝阳区人民法院管辖。</p>
<p>本协议最后更新：2026 年 1 月</p>`
