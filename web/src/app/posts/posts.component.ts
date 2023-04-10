import { Component, OnInit } from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { User } from '../shared/models/user.model';
import { UserService } from '../shared/services/user.service';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css'],
  providers: [UserService]
})
export class PostsComponent implements OnInit {
  
  user!: User;
  posts = [
    { title: 'First Post', content: 'This is the first post.' },
    { title: 'Second Post', content: 'This is the second post.' },
    { title: 'Third Post', content: 'This is the third post.' }
  ];

  

  constructor(
    private breakpointObserver: BreakpointObserver,
    private userService: UserService) { }

  ngOnInit(): void {
    this.user = new User(window.Telegram.WebApp.initData);
    this.createUser(this.user);
  }

  onClick(): void {
    console.log("click", this.user)
    this.getUsers()
  }

  onIncCaviar(name: string, size: number): void {
      
  }

  getUsers(): void {
    this.userService.getUsers()
      .subscribe(data => (
        console.log(data)
      ));
  }

  createUser(user: User): void {
    this.userService.crateUser(user)
      .subscribe(data => (
        console.log(data)
      ));
  }

  isSmallScreen: Observable<boolean> = this.breakpointObserver.observe(Breakpoints.Small)
  .pipe(
    map(result => result.matches)
  );

}
