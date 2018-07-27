package main

import (
	"reflect"

	"github.com/go-macaron/captcha"
	"gopkg.in/macaron.v1"
)

func homeView(ctx *macaron.Context) {
	ctx.Data["Title"] = ""

	sup := &SignupForm{Type: 4}
	ctx.Data["Form"] = sup

	ctx.HTML(200, "home")
}

func kontaktView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Kontakt | "

	ctx.HTML(200, "kontakt")
}

func kontaktViewPost(ctx *macaron.Context, contact ContactForm, cpt *captcha.Captcha) {
	ctx.Data["Title"] = "Contact | "

	em := &EmailMessage{}
	em.FromName = contact.Name
	em.FromEmail = contact.Email
	em.ToName = "Pragmatic"
	em.ToEmail = "anonutopia@protonmail.com"
	if len(contact.Subject) == 0 {
		em.Subject = "Contact Form Message"
	} else {
		em.Subject = contact.Subject
	}
	em.BodyHTML = contact.Message
	em.BodyText = contact.Message

	s := reflect.ValueOf(ctx.Data["Errors"])

	if s.Len() == 0 {
		if cpt.VerifyReq(ctx.Req) {
			err := sendEmail(em)
			if err != nil {
				ctx.Data["Form"] = contact
				ctx.Data["SendError"] = true
			} else {
				ctx.Data["Success"] = true
			}
		} else {
			ctx.Data["Form"] = contact
			ctx.Data["CaptchaError"] = true
		}
	} else {
		ctx.Data["Form"] = contact
	}

	ctx.HTML(200, "kontakt")
}

func view404(ctx *macaron.Context) {
	ctx.Data["URI"] = "/not-found/"
	ctx.Data["Title"] = "404 | "

	ctx.HTML(404, "404")
}
