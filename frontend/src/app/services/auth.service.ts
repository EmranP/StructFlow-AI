import { HttpClient } from '@angular/common/http'
import { Injectable, signal } from '@angular/core'
import { Observable, tap } from 'rxjs'
import { AuthResponse, GetID } from '../models'

@Injectable({ providedIn: 'root' })
export class AuthService {
	private apiUrl = 'http://localhost:3000/api'
	isLoggedIn = signal<boolean>(this.hasToken())

	constructor(private http: HttpClient) {}

	private hasToken(): boolean {
		return !!localStorage.getItem('token')
	}

	getToken(): string | null {
		return localStorage.getItem('token')
	}

	getMe(): Observable<GetID> {
		return this.http.get<GetID>(`${this.apiUrl}/auth/me`)
	}

	register(email: string, password: string): Observable<AuthResponse> {
		return this.http
			.post<AuthResponse>(`${this.apiUrl}/auth/register`, { email, password })
			.pipe(
				tap(res => {
					localStorage.setItem('token', res.accessToken)
					this.isLoggedIn.set(true)
				})
			)
	}

	login(email: string, password: string): Observable<AuthResponse> {
		return this.http
			.post<AuthResponse>(`${this.apiUrl}/auth/login`, { email, password })
			.pipe(
				tap(res => {
					localStorage.setItem('token', res.accessToken)
					this.isLoggedIn.set(true)
				})
			)
	}

	logout(): void {
		localStorage.removeItem('token')
		this.isLoggedIn.set(false)
	}
}
