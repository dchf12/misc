export default {
  async email(message, env, ctx) {
    await message.forward("example@example.com");
    const SLACK_WEBHOOK_URL = "https://hooks.slack.com/services/[credential]";

    const data = {
      username: message.from,
      channel: "#mail",
      text: `<!channel> ${message.from}からのメール`,
      attachments: [
        {
          fallback: `${message.from} → ${message.to}`,
          color: "#0000D0",
          fields: [
            {
              title: "送信元",
              value: message.from,
            },
            {
              title: "送信先",
              value: message.to,
            },
            {
              title: "件名",
              value: message.headers.get("subject") ?? "件名なし",
            },
          ],
        },
      ],
    };

    const result = await fetch(SLACK_WEBHOOK_URL, {
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      },
      method: "POST",
      body: JSON.stringify(data),
    });
    console.log(result);
  },
};
