const axios = require('axios');

const url = 'http://localhost:8080/api';

const job = {
    label: "some_descriptive_label",
    tasks: [
        {
            id: "TaskNumber-0-36b8d41a-8611-4468-93ee-40f4140c7555",
            spec: {
                image: "ubuntu",
                requirements: {
                    DockerRequirements: "DockerMemory == 1024 && DockerCPUWeight == 512"
                }
            },
            commands: [
                "sleep 10",
                "sleep 20"
            ]
        }
    ]
};

axios.post(`${url}/queues`, {
    "name": "awesome_name"
}).then((response) => {
    const id = response.data.id;
    for (let i = 0; i < 100; i++) {
        axios.post(`${url}/queues/${id}/jobs`, job).catch(() => {});
    }

    setTimeout(() => {
        for (let i = 0; i < 500; i++) {
            axios.post(`${url}/queues/${id}/jobs`, job).catch(() => {});
        }
    }, 10000);
});
