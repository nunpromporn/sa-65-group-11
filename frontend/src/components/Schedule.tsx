import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import { ScheduleInterface } from "../models/ISchedule";
import { format } from 'date-fns'

// const useStyles = makeStyles((theme: Theme) =>
//   createStyles({
//     container: {
//       marginTop: theme.spacing(2),
//     },
//     table: {
//       minWidth: 800,
//     },
//     tableSpace: {
//       marginTop: 20,
//     },
//   })
// );

function Schedules() {
  // const classes = useStyles();
  const [schedules, setSchedules] = useState<ScheduleInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getSchedules = async () => {
    fetch(`${apiUrl}/schedules`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
            setSchedules(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getSchedules();
  }, []);

  return (
    <div>
      <Container sx={{ marginTop: 2 }} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ตารางการทำงานของแพทย์
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/schedule/create"
              variant="contained"
              color="primary"
            >
              ตารางการทำงาน
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} sx={{ marginTop: 2 }}>
          <Table sx={{ minWidth: 800 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="2%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="18%">
                  ชื่อแพทย์
                </TableCell>
                <TableCell align="center" width="10%">
                  แผนก
                </TableCell>
                <TableCell align="center" width="15%">
                  สถานที่
                </TableCell>
                <TableCell align="center" width="20%">
                  ห้อง
                </TableCell>
                <TableCell align="center" width="30%">
                  เวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {schedules.map((item: ScheduleInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Doctor.Name}</TableCell>
                  <TableCell align="center">{item.Department.Name}</TableCell>
                  <TableCell align="center">{item.Location.Name}</TableCell>
                  <TableCell align="center">{item.Room.Name}</TableCell>
                  <TableCell align="center">{format((new Date(item.ScheduleTime)), 'dd MMMM yyyy hh:mm')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Schedules;