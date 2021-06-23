import { NgModule } from '@angular/core';
import { AvatarComponent } from './avatar.component';

const COMPONENTS = [AvatarComponent];

@NgModule({
    declarations: COMPONENTS,
    exports: COMPONENTS
})
export class AvatarModule {}
