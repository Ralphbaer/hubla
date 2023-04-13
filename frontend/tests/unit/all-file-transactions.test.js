import { fetchAllTransactions } from '../../js/all-file-transactions';

describe('fetchAllTransactions', () => {
    it('returns an array of transactions', async () => {
        const mockData = [
            {
                "id": "047439c3-7a8e-42e2-9849-0bdf4f177009",
                "t_type": 1,
                "t_date": "2022-02-01T23:35:43Z",
                "product_id": "75f8df81-e5e1-4d77-b65e-76a0a2eaf167",
                "amount": "155000",
                "seller_id": "ea09d1de-e97f-4a85-843e-c6260d8541e5",
                "created_at": "2023-04-12T08:11:44.20286Z"
            },
            {
                "id": "1983c7b4-320b-48a0-bacc-e5fde4861f2d",
                "t_type": 4,
                "t_date": "2022-02-03T17:23:37Z",
                "product_id": "75f8df81-e5e1-4d77-b65e-76a0a2eaf167",
                "amount": "50000",
                "seller_id": "427712c7-0384-4cf2-85eb-02c6c3cebfdb",
                "created_at": "2023-04-12T08:11:44.21526Z"
            },
        ];

        jest.mock('../../js/all-file-transactions', () => ({
            fetchAllTransactions: jest.fn(() => mockData),
        }));

        const transactions = await fetchAllTransactions();
        expect(Array.isArray(transactions)).toBe(true);
        expect(transactions.length).toBe(mockData.length);
    });

    it('throws an error when the response is not ok', async () => {
        jest.mock('../../js/all-file-transactions', () => ({
            fetchAllTransactions: jest.fn(() => {
                throw new Error('Unable to fetch transactions');
            }),
        }));

        await expect(fetchAllTransactions()).rejects.toThrow('Unable to fetch transactions');
    });
});
