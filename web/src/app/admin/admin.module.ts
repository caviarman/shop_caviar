import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { RouterModule } from "@angular/router";
import { LoginPageComponent } from "./login-page/login-page.component";
import { AdminLayoutComponent } from './shared/components/admin-layout/admin-layout.component';
import { DashboardPageComponent } from './dashboard-page/dashboard-page.component';
import { CreatePageComponent } from './create-page/create-page.component';
import { EditPageComponent } from './edit-page/edit-page.component';

@NgModule({
    imports: [
        CommonModule,
        RouterModule.forChild([
            {path: '', component: AdminLayoutComponent, children: [
                {path: '', redirectTo: '/admin/login', pathMatch: 'full'},
                {path: 'login', component: LoginPageComponent},
                {path: 'dashboard', component: DashboardPageComponent},
                {path: 'create', component: CreatePageComponent},
                {path: 'post/:id/edit', component: EditPageComponent}
            ]}
        ])
    ],
    exports: [
        RouterModule
    ],
    declarations: [
      AdminLayoutComponent,
      DashboardPageComponent,
      CreatePageComponent,
      EditPageComponent
    ]
})

export class AdminModule {


}