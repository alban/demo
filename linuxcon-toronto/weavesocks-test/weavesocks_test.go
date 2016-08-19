package potato_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("firefox"))
		Expect(err).NotTo(HaveOccurred())
		page.SetImplicitWait(10000)
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should give user feedback", func() {
		By("go to homepage", func() {
			Expect(page.Navigate("http://localhost:9101/")).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:9101/"))
		})

		By("login", func() {
			Eventually(page.Find("#login"), "10s").Should(BeFound())
			Eventually(page.Find("#login"), "10s").Should(HaveText("Login"))

			// Give some time to follow the demo
			time.Sleep(2 * time.Second)

			Expect(page.Find("#login").FindByLink("Login").Click()).To(Succeed())
			Expect(page.Find("#username-modal").Fill("user")).To(Succeed())
			Expect(page.Find("#password-modal").Fill("password")).To(Succeed())
			Expect(page.Find("#password-modal").Fill("password")).To(Succeed())
			time.Sleep(time.Second)
			Expect(page.Find(".fa-sign-in").Click()).To(Succeed())

			Expect(page).To(HavePopupText("Logged in as user"))
			Expect(page.ConfirmPopup()).To(Succeed())

			Eventually(page.Find("#howdy"), "10s").Should(HaveText("Logged in as User Name"))

		})

		By("go to catalogue", func() {
			Expect(page.FindByLink("Catalogue").Click()).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:9101/category.html"))
		})

		By("add item in the cart", func() {
			Expect(page.First(".product").Find(".fa-shopping-cart").Click()).To(Succeed())
			Eventually(page.Find("#numItemsInCart"), "10s").ShouldNot(HaveText("0 item(s) in cart"))
			time.Sleep(time.Second)
		})

		By("go to cart", func() {
			Expect(page.Find("#numItemsInCart").Click()).To(Succeed())
			Eventually(page.Find("#basket"), "10s").Should(MatchText("Shopping cart"))
		})

		// Give some time to follow the demo
		time.Sleep(5 * time.Second)
	})
})
