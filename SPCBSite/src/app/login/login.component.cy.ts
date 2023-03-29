import { LoginComponent } from './login.component'
import { HttpClientModule, HttpHandler } from '@angular/common/http';

describe('LoginComponent', () => {
    it('mounts', () => {
      cy.mount(LoginComponent)

    })
})