import { SignupComponent } from './signup.component'
import { HttpClientModule, HttpHandler } from '@angular/common/http';

describe('SignupComponent', () => {
    it('mounts', () => {
      cy.mount(SignupComponent)

    })
})