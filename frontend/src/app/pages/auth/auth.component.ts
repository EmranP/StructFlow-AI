import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss']
})
export class AuthComponent {
  auth = inject(AuthService);
  router = inject(Router);

  mode: 'login' | 'register' = 'login';
  email = '';
  password = '';
  loading = false;
  error = '';

  toggle() {
    this.mode = this.mode === 'login' ? 'register' : 'login';
    this.error = '';
  }

  submit() {
    if (!this.email || !this.password) { this.error = 'Fill all fields'; return; }
    this.loading = true;
    this.error = '';

    const obs = this.mode === 'login'
      ? this.auth.login(this.email, this.password)
      : this.auth.register(this.email, this.password);

    obs.subscribe({
      next: () => this.router.navigate(['/projects']),
      error: (e) => {
        this.loading = false;
        this.error = e?.error?.message || 'Something went wrong';
      }
    });
  }
}
