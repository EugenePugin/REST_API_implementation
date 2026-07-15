import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "@typespec/openapi"; // Добавили новый импорт

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI; // Подключили пространство имен OpenAPI

// Исправлено: разделяем title и version по разным декораторам согласно стандарту 1.0+
@service(#{
  title: "E-Commerce REST API Demo",
})
@info(#{
  version: "1.0.0",
})
@server("http://localhost:8080", "Локальный сервер разработки")
namespace ECommerceService;

// --- 1. Описание схем данных (моделей) ---

model Product {
  @doc("Идентификатор товара")
  id: string;
  @doc("Название товара")
  name: string;
}

model User {
  @doc("Идентификатор пользователя")
  id: string;
  @doc("Имя пользователя")
  name: string;
}

model UserOrders {
  @doc("Идентификатор пользователя, сделавшего заказ")
  user_id: string;
  @doc("Список названий заказов")
  orders: string[];
}

// Стандартная модель для ошибки 401
model UnauthorizedError {
  @statusCode statusCode: 401;
  error: "Unauthorized";
}

// --- 2. Описание схемы авторизации (через заголовок Authorization) ---
model CustomAuth is ApiKeyAuth<ApiKeyLocation.header, "Authorization">;

// --- 3. Описание эндпоинтов (Роутинг) ---

@route("/api/v1")
namespace ApiV1 {

  // Саброутер для ТОВАРОВ (Публичный)
  @route("/products")
  @tag("Products")
  namespace Products {
    @doc("Получить список всех товаров")
    @get op list(): string[];

    @doc("Получить информацию о товаре по его ID")
    @get op get(@path id: string): Product;
  }

  // Саброутер для ПОЛЬЗОВАТЕЛЕЙ (Приватный)
  @route("/users")
  @tag("Users")
  @useAuth(CustomAuth) // Применяем авторизацию ко всей группе /users
  namespace Users {
    @doc("Получить список всех пользователей")
    @get op list(): string[] | UnauthorizedError;

    @doc("Получить профиль пользователя по его ID")
    @get op get(@path id: string): User | UnauthorizedError;

    // Вложенный саброутер для ЗАКАЗОВ пользователя
    @route("/{id}/orders")
    @tag("Orders")
    @doc("Получить заказы конкретного пользователя")
    @get op listOrders(@path id: string): UserOrders | UnauthorizedError;
  }
}