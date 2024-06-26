// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/FACorreiaa/glasses-management-platform/app/models"
)

func RegisterPage(register models.RegisterPage) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"w-full bg-white\"><div class=\"mx-auto max-w-7xl\"><div class=\"flex flex-col lg:flex-row\"><div class=\"relative w-full bg-cover lg:w-6/12 xl:w-7/12 bg-gradient-to-r from-white via-white to-gray-100\"><div class=\"relative flex flex-col items-center justify-center w-full h-full px-10 my-20 lg:px-16 lg:my-0\"><div class=\"flex flex-col items-start space-y-8 tracking-tight lg:max-w-3xl\"><div class=\"relative\"><p class=\"mb-2 font-medium text-gray-700 uppercase\">Work smarter</p><h2 class=\"text-5xl font-bold text-gray-900 xl:text-6xl\">Features to help you work smarter</h2></div><p class=\"text-2xl text-gray-700\">We've created a simple formula to follow in order to gain more out of your business and your application.</p><a href=\"#_\" class=\"inline-block px-8 py-5 text-xl font-medium text-center text-white transition duration-200 bg-blue-600 rounded-lg hover:bg-blue-700 ease\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\">Get Started Today</a></div></div></div><div class=\"w-full bg-white lg:w-6/12 xl:w-5/12\"><div class=\"flex flex-col items-start justify-start w-full h-full p-10 lg:p-16 xl:p-24\"><h4 class=\"w-full text-3xl font-bold\">Signup</h4><p class=\"text-lg text-gray-500\">or, if you have an account you can <a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.URL("/login")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">sign in</a></p><div class=\"relative w-full mt-10 space-y-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if register.Errors != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"error-messages text-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, err := range register.Errors {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(err)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/view/user/register.templ`, Line: 31, Col: 19}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form method=\"post\"><fieldset><fieldset class=\"max-w-lg mx-auto mt-2\"><label for=\"username\" class=\"block text-sm font-medium text-gray-900 dark:text-white\">Username</label> <input class=\"block w-full px-4 py-4 mt-2 text-xl placeholder-gray-400 bg-gray-200 rounded-lg focus:outline-none focus:ring-4 focus:ring-blue-600 focus:ring-opacity-50\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\" type=\"text\" placeholder=\"Username\" required name=\"username\" autocomplete=\"username\" id=\"username\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(register.Values["Username"])
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/view/user/register.templ`, Line: 49, Col: 46}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></fieldset><fieldset class=\"max-w-lg mx-auto mt-2\"><label for=\"email\" class=\"block text-sm font-medium text-gray-900 dark:text-white\">Email</label> <input class=\"block w-full px-4 py-4 mt-2 text-xl placeholder-gray-400 bg-gray-200 rounded-lg focus:outline-none focus:ring-4 focus:ring-blue-600 focus:ring-opacity-50\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\" type=\"email\" placeholder=\"Email\" required name=\"email\" id=\"email\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(register.Values["Email"])
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/view/user/register.templ`, Line: 63, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></fieldset><fieldset class=\"max-w-lg mx-auto mt-2\"><label for=\"passoword\" class=\"block text-sm font-medium text-gray-900 dark:text-white\">Password</label> <input class=\"block w-full px-4 py-4 mt-2 text-xl placeholder-gray-400 bg-gray-200 rounded-lg focus:outline-none focus:ring-4 focus:ring-blue-600 focus:ring-opacity-50\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\" type=\"password\" placeholder=\"Password\" required name=\"password\" autocomplete=\"new-password\"></fieldset><fieldset class=\"max-w-lg mx-auto mt-2\"><label for=\"password\" class=\"block text-sm font-medium text-gray-900 dark:text-white\">Confirm Password</label> <input class=\"block w-full px-4 py-4 mt-2 text-xl placeholder-gray-400 bg-gray-200 rounded-lg focus:outline-none focus:ring-4 focus:ring-blue-600 focus:ring-opacity-50\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\" type=\"password\" placeholder=\"Confirm Password\" required name=\"password_confirm\" autocomplete=\"new-password\"></fieldset></fieldset><div class=\"flex items-center justify-center max-w-lg mx-auto mt-2\"><button type=\"submit\" class=\"focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800\">Signup</button></div></form></div></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
