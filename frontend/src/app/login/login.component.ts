import { Component, EventEmitter, OnInit, Output } from '@angular/core'
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { HeaderstateService } from '../services/headerstate.service';
import { TokenServiceService } from '../services/token-service.service';
import { UserserviceService } from '../services/userservice.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  constructor(
    private route: ActivatedRoute, 
    private router: Router,
     private userService: UserserviceService,
      public headerState: HeaderstateService,
       private tokenService: TokenServiceService) { }

  profileForm = new FormGroup({
    email: new FormControl('', [
      Validators.required, Validators.pattern("^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,4}$")]),
    password: new FormControl('', [Validators.required, Validators.minLength(1)]),
  });
  // constructor() { }

  get email() { return this.profileForm.get('email'); }
  get password() { return this.profileForm.get('password'); }

  ngOnInit(): void {
  }

  onSubmit() {
    console.log(this.profileForm.get('email')?.value);
    console.log(this.profileForm.get('password')?.value);

    this.userService.login2(this.profileForm.get('email')?.value, this.profileForm.get('password')?.value).subscribe(data => {
      console.log('Log in!')
      if (data != null) {

        this.tokenService.setToken(data.token)
        this.router.navigate(['table', { isLogged: true, email: this.profileForm.get('email')?.value }])
      }
    })
  }

}
