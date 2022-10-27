import { DoctorInterface } from "./IDoctor";
import { LocationInterface } from "./ILocation";
import { DepartmentInterface } from "./IDepartment";
import { AuthoritieInterface } from "./IAuthoritie";
import { RoomInterface } from "./IRoom";
import { UserInterface } from "./IUser";

export interface ScheduleInterface {

    ID: number,
    ScheduleTime:       Date,


    UserID:       number,                                                                           //////////********** */
    User:         UserInterface,                                                    ////////////********** */


    // AuthoritieID:       number,
    // Authoritie:         AuthoritieInterface,

    DepartmentID:       number,
    Department:         DepartmentInterface,
    
    DoctorID:           number,
    Doctor:             DoctorInterface,

    LocationID:         number,
    Location:           LocationInterface,

    RoomID:             number,
    Room:               RoomInterface,
   
   }