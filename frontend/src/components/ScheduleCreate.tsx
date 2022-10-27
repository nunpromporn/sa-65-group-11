import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import TextField from '@mui/material/TextField';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';

import { DoctorInterface } from "../models/IDoctor";
import { DepartmentInterface } from "../models/IDepartment";
import { LocationInterface } from "../models/ILocation";
import { RoomInterface } from "../models/IRoom";
//import { AuthoritieInterface } from "../models/IAuthoritie";
import { ScheduleInterface } from "../models/ISchedule";

import { UserInterface } from "../models/IUser";                                  /////********* */

import { InputLabel } from "@mui/material";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";

// const Alert = (props: AlertProps) => {
//   return <MuiAlert elevation={6} variant="filled" {...props} />;
// };

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />
})

// const useStyles = makeStyles((theme: Theme) =>
//   createStyles({
//     root: {
//       flexGrow: 1,
//     },
//     container: {
//       marginTop: theme.spacing(2),
//     },
//     paper: {
//       padding: theme.spacing(2),
//       color: theme.palette.text.secondary,
//     },
//   })
// );

function BookingCreate() {
  // const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(null);
  //const [authorities, setAuthorities] = useState<AuthoritieInterface>(); //map

  const [users, setUsers] = useState<UserInterface>();                  ////////******* */


  const [departments, setDepartments] = useState<DepartmentInterface[]>([]);
  const [doctors, setDoctors] = useState<DoctorInterface[]>([]);
  const [locations, setLocations] = useState<LocationInterface[]>([]);
  const [rooms, setRooms] = useState<RoomInterface[]>([]);
  const [schedules, setSchedules] = useState<Partial<ScheduleInterface>>({});
  
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: SelectChangeEvent<number>
  ) => {
    const name = event.target.name as keyof typeof schedules;
    setSchedules({
      ...schedules,
      [name]: event.target.value,
    });
  };

  const  handledoctorChange  = (
    event: React.ChangeEvent<{ name?: string; value: any }>
  ) => {   
    const name = event.target.name as keyof typeof schedules;
    setSchedules({
      ...schedules,
      [name]: Number(event.target.value),
    });   
  };



  /*
  const getAuthorities = async () => {
    const uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/authoritie/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {                                                                                
        schedules.AuthoritieID = res.data.ID
        if (res.data) {
            setAuthorities(res.data);
        } else {
          console.log("else");
        }
      });
  };
*/

  const getUsers = async () => {
    const uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/user/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        schedules.UserID = res.data.ID
        if (res.data) {
            setUsers(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getDepartment = async () => {
    fetch(`${apiUrl}/departments`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDepartments(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getDoctor = async () => {
    fetch(`${apiUrl}/doctors`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDoctors(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getLocation = async () => {
    fetch(`${apiUrl}/locations`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setLocations(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getRoom = async () => {
    fetch(`${apiUrl}/rooms`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRooms(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {

    getUsers();             ////////////
    //getAuthorities();
    getDoctor();
    getDepartment();
    getLocation();
    getRoom();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {

        UserID: convertType(schedules.UserID),                        //////////////////
        // AuthoritieID: convertType(schedules.AuthoritieID),
        DepartmentID: convertType(schedules.DepartmentID),
        DoctorID: convertType(schedules.DoctorID),
        LocationID: convertType(schedules.LocationID),
        RoomID: convertType(schedules.RoomID),
        ScheduleTime: selectedDate,  

    };

    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/schedules`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true)
          setErrorMessage("")
        } else {
          console.log("บันทึกไม่ได้")
          setError(true)
          setErrorMessage(res.error)
        }
      });
  }

  return (
    <Container sx={{ marginTop: 2 }} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ: {errorMessage}
        </Alert>
      </Snackbar>
      <Paper sx={{ padding: 2, color: "text.secondary" }}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h6"
              variant="h5"
              color="primary"
              gutterBottom
              
            >
              บันทึกตารางการทำงานของแพทย์

            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3}  sx={{ flexGrow: 1 }}>

            <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                <Select
                    native
                    disabled
                    value={schedules.UserID}                                      //////////
                    // label="ชื่อ - นามสกุล"
                    onChange={handleChange}
                    // inputProps={{
                    // name: "PatientID",
                    // }}
                > 

                    <option value={users?.ID} key={users?.ID} >                 
                    {users?.Name}                                                 
                    </option>
                    {/* {authorities.map((item: PatientInterface) => (
                    <option value={item.ID} key={item.ID}>
                        {item.Name}
                    </option>
                    ))} */}
                    
                </Select>
                </FormControl>
            </Grid>

            <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                <InputLabel id="DoctorID">เลือกแพทย์</InputLabel>
                <Select
                    native
                    value={schedules.DoctorID}
                    label="เลือกแพทย์"
                    onChange={handleChange}
                    inputProps={{
                    name: "DoctorID",
                    }}
                >
                    <option aria-label="None" value="">
                    </option>
                    {doctors.map((item: DoctorInterface) => (
                    <option value={item.ID} key={item.ID}>
                        {item.Name}
                    </option>
                    ))}
                </Select>
                </FormControl>
            </Grid>

            <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                <InputLabel id="DepartmentID">เลือกแผนกทางการแพทย์</InputLabel>
                <Select
                    native
                    value={schedules.DepartmentID}
                    label="เลือกแผนกทางการแพทย์"
                    onChange={handleChange}
                    inputProps={{
                    name: "DepartmentID",
                    }}
                >
                    <option aria-label="None" value="">
                    </option>
                    {departments.map((item: DepartmentInterface) => (
                    <option value={item.ID} key={item.ID}>
                        {item.Name}
                    </option>
                    ))}
                </Select>
                </FormControl>
            </Grid>

            <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                <InputLabel id="LocationID">เลือกสถานที่</InputLabel>
                <Select
                    native
                    value={schedules.LocationID}
                    label="เลือกสถานที่"
                    onChange={handleChange}
                    inputProps={{
                    name: "LocationID",
                    }}
                >
                    <option aria-label="None" value="">
                    </option>
                    {locations.map((item: LocationInterface) => (
                    <option value={item.ID} key={item.ID}>
                        {item.Name}
                    </option>
                    ))}
                </Select>
                </FormControl>
            </Grid>

            <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                <InputLabel id="RoomID">เลือกห้อง</InputLabel>
                <Select
                    native
                    value={schedules.RoomID}
                    label="เลือกห้อง"
                    onChange={handleChange}
                    inputProps={{
                    name: "RoomID",
                    }}
                >
                    <option aria-label="None" value="">
                    </option>
                    {rooms.map((item: RoomInterface) => (
                    <option value={item.ID} key={item.ID}>
                        {item.Name}
                    </option>
                    ))}
                </Select>
                </FormControl>
            </Grid>

            <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DateTimePicker
                  label="เลือกวันเวลา"
                  value={selectedDate}
                  onChange={(newValue) => setSelectedDate(newValue)}
                  minDate={(new Date('31-12-2022T09:00'))}
                  renderInput={(params) => 
                  <TextField {...params} />}
                />
                </LocalizationProvider>
              </FormControl>
            </Grid>
          
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/schedules"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default BookingCreate;