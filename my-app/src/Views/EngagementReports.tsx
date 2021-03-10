import { Header } from '../Components/Header';
import { BarChart, Bar, XAxis, YAxis, Tooltip } from 'recharts';

export type EngagementReportBarChartData = {
    name: string;
    count: number;
}

export function EngagementReports() {
    const testData = [
        {
          name: 'Topic 1',
          count: 250
        },
        {
          name: 'Topic 2',
          count: 170
        },
        {
          name: 'Topic 3',
          count: 192
        },
        {
          name: 'Topic 4',
          count: 32
        },
        {
          name: 'Topic 5',
          count: 45
        }
      ];

    testData.sort((a, b) => {
        if (a.count > b.count) {
            return -1;
        } else if (a.count < b.count) {
            return 1;
        } else {
            return 0;
        }
    });

    return (
        <div className="dashboard sub-dashboard">
            <Header title="Engagement Reports"></Header>
            <h2>Issues/Inquiries</h2>
            <BarChart
                width={700}
                height={500}
                data={testData}
                margin={{
                    top: 5,
                    right: 30,
                    left: 20,
                    bottom: 5,
                }}
                >
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Bar dataKey="count" fill="#855CF8" />
            </BarChart>
        </div>
    );
}