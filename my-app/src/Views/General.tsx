import { SubHeader } from '../Components/SubHeader';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';

// TODO: will data be pre-sorted on back-end?
function getSubDashboardData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "SPD - Reform", value: 123});
    data.push({name: "COVID-19 - Stimulus", value: 119});
    data.push({name: "Homelessness- Shelter", value: 77});
    data.push({name: "Investments in BIPOC Communities", value: 62});
    data.push({name: "SPD - Accountability", value: 36});
    data.push({name: "Other", value: 52});
    return data as Array<SubDashboardData>;
}

export function General() {
    const test_data = getSubDashboardData();

    return (
        <div>
            <SubHeader title="General Inquiries"></SubHeader>
            <SubDashboard title="TOPIC" data={test_data}></SubDashboard>
        </div>
    );
}