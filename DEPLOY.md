# Deployment Guide - PayX Payment Gateway

This guide will help you deploy PayX to various platforms.

## Prerequisites

- A PostgreSQL database (Supabase, AWS RDS, or any PostgreSQL provider)
- Midtrans account credentials
- Git repository access

## Deployment Options

### Option 1: Render.com

1. **Setup Database (Supabase)**
   - Create account at [Supabase](https://supabase.com)
   - Create a new project
   - Navigate to Settings > Database
   - Copy the connection string
   - Format: `postgresql://postgres:[password]@db.[project-ref].supabase.co:5432/postgres`

2. **Deploy to Render**
   - Fork or push this repository to GitHub
   - Sign up at [Render](https://render.com)
   - Create a new Web Service
   - Connect your GitHub repository
   - Render will auto-detect configuration from `render.yaml`

3. **Configure Environment Variables**
   ```
   PORT=10000
   DB_HOST=db.your-project-ref.supabase.co
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your-supabase-password
   DB_NAME=postgres
   JWT_SECRET_KEY=your-super-secret-jwt-key-min-32-chars
   JWT_EXPIRATION_HOURS=24h
   MIDTRANS_SERVER_KEY=your-midtrans-server-key
   MIDTRANS_CLIENT_KEY=your-midtrans-client-key
   MIDTRANS_ENVIRONMENT=sandbox
   ```

### Option 2: Heroku

1. **Install Heroku CLI**
   ```bash
   heroku login
   ```

2. **Create Heroku App**
   ```bash
   heroku create your-app-name
   ```

3. **Add PostgreSQL Addon**
   ```bash
   heroku addons:create heroku-postgresql:hobby-dev
   ```

4. **Set Environment Variables**
   ```bash
   heroku config:set JWT_SECRET_KEY=your-secret-key
   heroku config:set MIDTRANS_SERVER_KEY=your-server-key
   heroku config:set MIDTRANS_CLIENT_KEY=your-client-key
   heroku config:set MIDTRANS_ENVIRONMENT=sandbox
   ```

5. **Deploy**
   ```bash
   git push heroku main
   ```

### Option 3: Docker

1. **Build Docker Image**
   ```bash
   docker build -t payx .
   ```

2. **Run Container**
   ```bash
   docker run -d \
     -p 8080:8080 \
     --env-file .env \
     --name payx \
     payx
   ```

### Option 4: AWS/GCP/Azure

1. **Setup PostgreSQL Database**
   - Create a managed PostgreSQL instance
   - Note connection details

2. **Deploy Application**
   - Use container services (ECS, Cloud Run, Container Instances)
   - Configure environment variables
   - Set up load balancer if needed

## Post-Deployment Checklist

- [ ] Verify database connection
- [ ] Test authentication endpoints
- [ ] Configure webhook URL in Midtrans dashboard
- [ ] Test payment creation flow
- [ ] Monitor application logs
- [ ] Set up SSL/TLS certificates
- [ ] Configure CORS if needed
- [ ] Set up monitoring and alerts

## Webhook Configuration

Configure the webhook URL in your Midtrans dashboard:
```
https://your-domain.com/api/v1/payments/notification
```

## Troubleshooting

### Database Connection Issues
- Verify database credentials
- Check network connectivity
- Ensure database allows connections from your deployment IP

### Payment Processing Issues
- Verify Midtrans credentials
- Check environment (sandbox vs production)
- Review Midtrans dashboard for transaction logs

### Authentication Issues
- Verify JWT_SECRET_KEY is set
- Check token expiration settings
- Review request headers

## Monitoring

Monitor your deployment:
- Application logs
- Database performance
- API response times
- Error rates
- Transaction success rates

## Security Best Practices

- Use strong JWT secret keys
- Enable HTTPS/SSL
- Regularly rotate credentials
- Monitor for suspicious activity
- Keep dependencies updated
- Use environment variables for secrets
- Implement rate limiting
- Set up proper CORS policies
